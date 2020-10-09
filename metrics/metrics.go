package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Validate - Prometheus Histogram
var Validate = promauto.NewHistogram(prometheus.HistogramOpts{
	Name: "Validate",
	Help: "Metric for GET calls to /validate",
})

// ValidateSCRandomTimeMethod - Prometheus Histogram
var ValidateSCRandomTimeMethod = promauto.NewHistogram(prometheus.HistogramOpts{
	Name: "Validate_SC_RandomTimeMethod",
	Help: "Amount of time for RandomTimeMethod() to execute",
})

// PromTimer - Helper method
func PromTimer(startTime time.Time, histogram prometheus.Histogram) {
	totalDuration := time.Since(startTime)
	seconds := float64(totalDuration) / float64(time.Second)
	histogram.Observe(seconds)
}
