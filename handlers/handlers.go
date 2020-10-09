package handlers

import (
	"fmt"
	"net/http"
)

// Hola handler
func Hola(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola")
}

// HelloServer handler
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
