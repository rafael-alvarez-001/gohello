package main

import (
	"gohello/metrics"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	metrics.HelloRalph()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
