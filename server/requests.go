package main

import (
	"io"
	"log"
	"net/http"
)

// Структура, ответственная за части запроса
type RequestParts struct {
	// Тип запроса
	// (GET, POST, PUT, PATCH, DELETE)
	Type   string
	URL    string
	Header io.Reader
	Body   io.Reader
}

// Сборка частей запроса в полноценный запрос
func (r *RequestParts) CreateRequest() *http.Request {
	var req *http.Request
	var err error
	req, err = http.NewRequest(r.Type, r.URL, r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return req
}

// Отправка запроса
func SendRequest(req *http.Request) *http.Response {
	var resp *http.Response
	var err error

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	return resp
}

// Чтение ответа
func ReadResponse(resp *http.Response) string {
	defer resp.Body.Close()
	var body []byte
	var err error

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(body)
}
