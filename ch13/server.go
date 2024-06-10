package spewg

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

const replicationHeader = "X-Replication-Request"

type CacheServer struct {
	cache    *Cache
	peers    []string
	hashRing *HashRing
	mu       sync.Mutex
}

func NewCacheServer(peers []string) *CacheServer {
	cs := &CacheServer{
		cache:    NewCache(10),
		peers:    peers,
		hashRing: NewHashRing(),
	}
	for _, peer := range peers {
		cs.hashRing.AddNode(Node{ID: peer, Addr: peer})
	}

	return cs
}

func (cs *CacheServer) SetHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	targetNode := cs.hashRing.GetNode(req.Key)
	if targetNode.Addr == "self" { // "self" indicates the current node
		log.Printf("Setting key %q with value %q on current node", req.Key, req.Value)
		cs.cache.Set(req.Key, req.Value, 1*time.Hour)
		if r.Header.Get(replicationHeader) == "" {
			go cs.replicateSet(req.Key, req.Value)
		}
	} else {
		log.Printf("Forwarding set request for key %q to node %q", req.Key, targetNode.Addr)
		cs.forwardRequest(targetNode, r)
	}
	w.WriteHeader(http.StatusOK)
}

func (cs *CacheServer) replicateSet(key, value string) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	req := struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}{
		Key:   key,
		Value: value,
	}

	data, _ := json.Marshal(req)

	for _, peer := range cs.peers {
		if peer != "self" {
			go func(peer string) {
				client := &http.Client{}
				req, err := http.NewRequest("POST", peer+"/set", bytes.NewReader(data))
				if err != nil {
					log.Printf("Failed to create replication request: %v", err)
					return
				}
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set(replicationHeader, "true")
				_, err = client.Do(req)
				if err != nil {
					log.Printf("Failed to replicate to peer %s: %v", peer, err)
				}
			}(peer)
		}
	}
}

func (cs *CacheServer) GetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	targetNode := cs.hashRing.GetNode(key)
	if targetNode.Addr == "self" {
		value, found := cs.cache.Get(key)
		if !found {
			http.NotFound(w, r)
			return
		}
		err := json.NewEncoder(w).Encode(map[string]string{"value": value})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		cs.forwardRequest(targetNode, r)
	}
}

func (cs *CacheServer) forwardRequest(targetNode Node, r *http.Request) {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 100, // Adjust based on your load
		},
		Timeout: 5 * time.Second, // Prevent requests from hanging indefinitely
	}

	// Create a new request based on the method
	var req *http.Request
	var err error

	if r.Method == http.MethodGet {
		// Forward GET request with query parameters
		getURL := fmt.Sprintf("%s%s?%s", targetNode.Addr, r.URL.Path, r.URL.RawQuery)
		req, err = http.NewRequest(r.Method, getURL, nil)
	} else if r.Method == http.MethodPost {
		// Forward POST request with body
		postURL := fmt.Sprintf("%s%s", targetNode.Addr, r.URL.Path)
		req, err = http.NewRequest(r.Method, postURL, r.Body)
	}

	if err != nil {
		log.Printf("Failed to create forward request: %v", err)
		return
	}

	// Copy the headers
	req.Header = r.Header

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		// Check for a "connection refused" error
		var urlErr *url.Error
		if errors.As(err, &urlErr) && urlErr.Err != nil {
			var opErr *net.OpError
			if errors.As(urlErr.Err, &opErr) && opErr.Op == "dial" {
				var sysErr *os.SyscallError
				if errors.As(opErr.Err, &sysErr) && sysErr.Syscall == "connect" {
					log.Printf("Connection refused to node %s: %v", targetNode.Addr, err)
					// Consider adding retry logic or node status checks here
					return
				}
			}
		}
		log.Printf("Failed to forward request to node %s: %v", targetNode.Addr, err)
		return
	}
	io.ReadAll(resp.Body)
}
