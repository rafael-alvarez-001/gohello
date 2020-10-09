package handlers

import (
	"fmt"
	"gohello/metrics"
	"net/http"
	"time"
)

// Hola handler
func Hola(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola")
}

// HelloServer handler
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

// ValidateHandler handler
func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	var lStart time.Time
	start := time.Now().UTC()

	time.Sleep(1 * time.Second)

	fmt.Fprintf(w, "Hola")

	lStart = time.Now().UTC()
	randomTimeMethod()
	metrics.PromTimer(lStart, metrics.ValidateSCRandomTimeMethod)

	metrics.PromTimer(start, metrics.Validate)

	fmt.Fprintf(w, "You've been validated!")

}

func randomTimeMethod() {
	time.Sleep(4 * time.Second)
}
