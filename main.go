package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Age  uint8
}

func home_page(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Go is super easy")
	bob := User{"Bob", 25}
	tmpl, _ := template.ParseFiles("views/home_page.html")
	tmpl.Execute(w, bob)
}

func handeRequest() {
	http.HandleFunc("/", home_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handeRequest()
}
