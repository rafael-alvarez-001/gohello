package main

import (
	"gohello/handlers"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func startMetrics() {
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":2112", nil)
}

func main() {

	startMetrics()

	http.HandleFunc("/hola", handlers.Hola)
	http.ListenAndServe(":8080", nil)
}
