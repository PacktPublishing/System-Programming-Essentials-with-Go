package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	spewg "github.com/alexrios/chapter13"
)

var port string
var peers string

func main() {
	flag.StringVar(&port, "port", ":8080", "HTTP server port")
	flag.StringVar(&peers, "peers", "", "Comma-separated list of peer addresses")

	flag.Parse()

	peerList := strings.Split(peers, ",")

	cs := spewg.NewCacheServer(append(peerList, "self"))
	http.HandleFunc("/set", cs.SetHandler)
	http.HandleFunc("/get", cs.GetHandler)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
