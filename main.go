package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Status  int
	Message string
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	data := response{http.StatusOK, "Hello, World"}

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	port := ":8080"
	http.HandleFunc("/", helloHandler)
	log.Printf("Starting server on %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
