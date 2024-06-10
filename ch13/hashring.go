package spewg

import (
	"crypto/sha1"
	"sort"
	"sync"
)

// Node represents a physical or virtual node in the system
type Node struct {
	ID   string
	Addr string
}

// HashRing represents the consistent hash ring
type HashRing struct {
	nodes  []Node
	hashes []uint32
	lock   sync.RWMutex
}

// NewHashRing creates a new HashRing
func NewHashRing() *HashRing {
	return &HashRing{}
}

// AddNode adds a node to the hash ring
func (h *HashRing) AddNode(node Node) {
	h.lock.Lock()
	defer h.lock.Unlock()
	hash := h.hash(node.ID)
	h.nodes = append(h.nodes, node)
	h.hashes = append(h.hashes, hash)
	sort.Slice(h.hashes, func(i, j int) bool {
		return h.hashes[i] < h.hashes[j]
	})
}

// RemoveNode removes a node from the hash ring
func (h *HashRing) RemoveNode(nodeID string) {
	h.lock.Lock()
	defer h.lock.Unlock()
	var index int
	var hash uint32
	for i, node := range h.nodes {
		if node.ID == nodeID {
			hash = h.hash(node.ID)
			index = i
			break
		}
	}
	h.nodes = append(h.nodes[:index], h.nodes[index+1:]...)
	for i, hsh := range h.hashes {
		if hsh == hash {
			h.hashes = append(h.hashes[:i], h.hashes[i+1:]...)
			break
		}
	}
}

// GetNode returns the node responsible for the given key
func (h *HashRing) GetNode(key string) Node {
	if len(h.nodes) == 0 {
		return Node{}
	}

	h.lock.RLock()
	defer h.lock.RUnlock()
	hash := h.hash(key)
	index := sort.Search(len(h.hashes), func(i int) bool {
		return h.hashes[i] >= hash
	})
	if index == len(h.hashes) {
		index = 0
	}
	return h.nodes[index]
}

func (h *HashRing) hash(key string) uint32 {
	hsh := sha1.New()
	hsh.Write([]byte(key))
	return h.bytesToUint32(hsh.Sum(nil))
}

func (h *HashRing) bytesToUint32(b []byte) uint32 {
	return (uint32(b[0]) << 24) | (uint32(b[1]) << 16) | (uint32(b[2]) << 8) | uint32(b[3])
}
