package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Loading index.html page
func indexPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/index.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "index", nil)
}

func handlers() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", indexPage)
}

func main() {
	fmt.Println("Starting server...")
	handlers()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Print(err.Error())
	}
}
