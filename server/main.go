package main

import (
	"encoding/json"
	"fmt"
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

type requestBody struct {
	Method  string            `json:"method"`
	ApiUrl  string            `json:"apiUrl"`
	Body    string            `json:"body"`
	Headers map[string]string `json:"headers"`
	Params  map[string]string `json:"params"`
}

func routeTest(w http.ResponseWriter, r *http.Request) {
	r.Header.Write(log.Writer())

	request_body := requestBody{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request_body); err != nil {
		// some error handling
		fmt.Println(err)
	}
	defer r.Body.Close()
	method := request_body.Method
	apiUrl := request_body.ApiUrl
	body := request_body.Body
	headers := request_body.Headers
	params := request_body.Params
	fmt.Println(method)
	fmt.Println(apiUrl)
	fmt.Println(body)
	fmt.Println(headers)
	fmt.Println(params)
}

func QueryHandler(w http.ResponseWriter, r *http.Request) {

	// http.ReadRequest(&bufio.Reader{r.URL})
	// var kal_govna []string
	// err := json.NewDecoder(r.Body).Decode(&kal_govna)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// fmt.Println()
	// fmt.Println(kal_govna)

	// kal_govna2 := r.RequestURI

	// fmt.Println()
	// fmt.Println(kal_govna2)
}

func main() {
	// /http.HandleFunc("/users", getUsers)

	fs := http.FileServer(http.Dir("../client/build"))
	http.Handle("/", fs)
	http.HandleFunc("/data", routeTest)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
