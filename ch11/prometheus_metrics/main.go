package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestsProcessed = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_processed",
			Help: "Total number of processed HTTP requests.",
		},
		[]string{"status_code"},
	)
)

func init() {
	prometheus.MustRegister(requestsProcessed)
}

func main() {
	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(50 * time.Millisecond)
		code := http.StatusOK
		if time.Now().Unix()%2 == 0 {
			code = http.StatusInternalServerError
		}
		requestsProcessed.WithLabelValues(fmt.Sprintf("%d", code)).Inc()
		w.WriteHeader(code)
		fmt.Fprintf(w, "Request processed.")
	})

	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}
