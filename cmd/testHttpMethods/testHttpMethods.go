package main

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/Yorik8877/gostman/scripts"
)

func main() {
	values := map[string]string{"foo1": "baz1"}
	jsonData, err := json.Marshal(values)
	if err != nil {
		panic(err)
	}

	var rp scripts.RequestParts = scripts.RequestParts{
		Type: "POST",
		URL:  "https://httpbin.org/post",
		Body: bytes.NewBuffer(jsonData),
	}
	// Полноценный реквест`
	req := rp.CreateRequest()

	// Отправка запроса и получение ответа
	resp := scripts.SendRequest(req)

	// Чтение ответа
	log.Println(scripts.ReadResponse(resp))
}
