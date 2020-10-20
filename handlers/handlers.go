package handlers

import (
	"fmt"
	"gohello/metrics"
	"net/http"
	"time"
)

// Pom handler
func Pom(w http.ResponseWriter, r *http.Request) {
	start := time.Now().UTC()

	fmt.Fprintf(w, "Pom baby")

	time.Sleep(1 * time.Second)

	metrics.PromTimer(start, metrics.PomEndpoint)
}

// Hola handler
func Hola(w http.ResponseWriter, r *http.Request) {
	start := time.Now().UTC()

	fmt.Fprintf(w, "Hola")

	time.Sleep(2 * time.Second)

	metrics.PromTimer(start, metrics.HolaEndpoint)
}

// ValidateHandler handler
func ValidateHandler(w http.ResponseWriter, r *http.Request) {

	var lStart time.Time
	start := time.Now().UTC()

	fmt.Fprintf(w, "Start Validation\n")

	lStart = time.Now().UTC()

	randomTimeMethod()
	metrics.PromTimer(lStart, metrics.ValidateSCRandomTimeMethod)

	fmt.Fprintf(w, "You've been validated!")

	metrics.PromTimer(start, metrics.ValidateEndpoint)

}

func randomTimeMethod() {
	time.Sleep(3 * time.Second)
	fmt.Println("This is io....")
}
