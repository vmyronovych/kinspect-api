package hello

import (
	"fmt"
	"net/http"
)

// HelloHandler handles requests to the "/hello" endpoint
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type to plain text
	w.Header().Set("Content-Type", "text/plain")
	// Respond with "Hello, World!"
	fmt.Fprintf(w, "Hello, World! Upd v6")
}

// SetupRoutes registers the "/hello" route for this feature
func SetupRoutes() {
	// Register the helloHandler for the "/hello" route
	http.HandleFunc("/hello", HelloHandler)
}

func init() {
	SetupRoutes() // Register routes for this handler
}
