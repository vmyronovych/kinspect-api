package main

import (
	"fmt"
	"log"
	"net/http"
	_ "vmyronovych/kinspect/handlers/hello"
)

func main() {
	// Start the server on port 8080
	fmt.Println("Server is listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
