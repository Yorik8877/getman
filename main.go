package main

import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!\n"))
}

// main
func main() {
	http.HandleFunc("/", Handler)
	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
