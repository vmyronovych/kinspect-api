package hello

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// HelloHandler handles requests to the "/hello" endpoint
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	errorCodeParam := r.URL.Query().Get("error_code")
	if errorCodeParam != "" {
		errorCode, err := strconv.Atoi(errorCodeParam)
		if err == nil {
			latency := time.Duration((rand.Intn(5) + 5)) * time.Second
			time.Sleep(latency)
			w.WriteHeader(errorCode)
			fmt.Fprintf(w, "failed!")
			return
		}
	}

	latency := time.Duration((rand.Intn(100) + 200)) * time.Millisecond
	time.Sleep(latency)
	w.WriteHeader(http.StatusOK) // 200 OK
	fmt.Fprintf(w, "success")
}

// SetupRoutes registers the "/hello" route for this feature
func SetupRoutes() {
	// Register the helloHandler for the "/hello" route
	http.HandleFunc("/hello", HelloHandler)
}

func init() {
	SetupRoutes() // Register routes for this handler
}
