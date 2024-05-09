package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Вызов index.html
func indexPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/index.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	tmpl.ExecuteTemplate(w, "index", nil)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", indexPage)
	http.ListenAndServe(":8080", nil)
}
