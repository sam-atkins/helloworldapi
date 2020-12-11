package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response struct is used for http body responses
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	data := Response{http.StatusOK, "Hello, World"}
	responseData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	data := Response{http.StatusOK, "Service is healthy"}
	responseData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}

// ServerMux provides http routing
func ServerMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/health", healthHandler)
	return mux
}

func main() {
	mux := ServerMux()
	port := ":8080"
	log.Printf("Starting server on %s\n", port)
	http.ListenAndServe(port, mux)
}
