package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRoot)               // Root endpoint
	http.HandleFunc("/api/message", handleMessage) // Correct route
	http.HandleFunc("/api/health", handleHealth)   // Health endpoint

	port := ":8080"
	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the SportLife API Server!"))
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{"status": "success", "message": "Server is healthy"}
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var reqBody struct {
			Message string `json:"message"`
		}

		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if reqBody.Message == "" {
			http.Error(w, `{"status":"fail", "message":"Missing 'message' field"}`, http.StatusBadRequest)
			return
		}

		response := map[string]string{"status": "success", "message": "Data received successfully"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
