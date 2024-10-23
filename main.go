package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response structure for the /hello endpoint
type Response struct {
	Message string `json:"message"`
}

// helloHandler returns a JSON response with a greeting message.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "Hello, World!"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Register the /hello endpoint
	http.HandleFunc("/hello", helloHandler)

	// Start the server on port 8080
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
