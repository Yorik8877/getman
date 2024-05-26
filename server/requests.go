package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

// Структура, ответственная за части запроса
type RequestParts struct {
	Method  string              `json:"method"`
	ApiUrl  string              `json:"apiUrl"`
	Body    string              `json:"body"`
	Headers map[string][]string `json:"headers"`
	Params  map[string]string   `json:"params"`
}

// Сборка частей запроса в полноценный запрос
func (r *RequestParts) CreateRequest() *http.Request {
	method := strings.ToUpper(r.Method)
	customUrl := r.ApiUrl

	if len(r.Params) > 0 {
		customUrl += "?"
		for key, value := range r.Params {
			customUrl += key + "=" + value + "&"
		}
	}
	req, err := http.NewRequest(method, customUrl, strings.NewReader(r.Body))
	if err != nil {
		log.Println(err)
	}
	return req
}

// Отправка запроса
func SendRequest(r *RequestParts) *http.Response {
	req := r.CreateRequest()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	return resp
}

// Чтение ответа
func ReadResponseBody(r *http.Response) string {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	return string(body)
}
