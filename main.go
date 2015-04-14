package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(handler)

	http.Handle("/", handler)

	http.ListenAndServe(":8080", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		get(w, r)
	case "POST":
		post(w, r)
	default:
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The result is %v", "888")
}

func post(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "SAVED result: %v", "888")
}
