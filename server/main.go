package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	request := RequestParts{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		log.Println(err)
	}
	defer r.Body.Close()
	apiResponse := SendRequest(&request)
	responseParts := CutResponse(apiResponse)
	json.NewEncoder(w).Encode(responseParts)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("../client/build")))
	http.HandleFunc("/data", QueryHandler)

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
