package main

import (
	"log"
	"net/http"
)

// type User struct {
// 	ID    int    `json:"id"`
// 	Name  string `json:"name"`
// 	Email string `json:"email"`
// }

// func getUsers(w http.ResponseWriter, r *http.Request) {
// 	users := []User{
// 		{ID: 1, Name: "Alice", Email: "alice@example.com"},
// 		{ID: 2, Name: "Bob", Email: "bob@example.com"},
// 	}

// 	json.NewEncoder(w).Encode(users)
// }

func main() {
	log.Println("Starting server...")

	fs := http.FileServer(http.Dir("client/build"))
	http.Handle("/", fs)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
