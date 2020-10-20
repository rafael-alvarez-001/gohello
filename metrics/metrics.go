package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// ValidateEndpoint - Prometheus Histogram
var ValidateEndpoint = promauto.NewHistogram(prometheus.HistogramOpts{
	Name:        "Validate_Endpoint",
	Help:        "Metric for GET calls to /validate",
	ConstLabels: map[string]string{"endpoint": "/validate", "result": "endpoint"},
})

// ValidateSCRandomTimeMethod - Prometheus Histogram
var ValidateSCRandomTimeMethod = promauto.NewHistogram(prometheus.HistogramOpts{
	Name: "Validate_SC_RandomTimeMethod",
	Help: "Amount of time for RandomTimeMethod() to execute",
})

// HolaEndpoint - Prometheus Histogram
var HolaEndpoint = promauto.NewHistogram(prometheus.HistogramOpts{
	Name:        "Hola_Endpoint",
	Help:        "Metric for GET calls to /hola",
	ConstLabels: map[string]string{"endpoint": "/hola", "result": "endpoint"},
})

// PomEndpoint - Prometheus Histogram
var PomEndpoint = promauto.NewHistogram(prometheus.HistogramOpts{
	Name:        "Pom_Endpoint",
	Help:        "Metric for GET calls to /pom",
	ConstLabels: map[string]string{"endpoint": "/pom", "result": "endpoint"},
})

// PromTimer - Helper method
func PromTimer(startTime time.Time, histogram prometheus.Histogram) {
	totalDuration := time.Since(startTime)
	seconds := float64(totalDuration) / float64(time.Second)
	histogram.Observe(seconds)
}
