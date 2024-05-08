package scripts

import (
	"io"
	"log"
	"net/http"
)

type RequestParts struct {
	Type string
	URL  string
	Body io.Reader
}

func (r *RequestParts) CreateRequest() *http.Request {
	var req *http.Request
	var err error
	req, err = http.NewRequest(r.Type, r.URL, r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return req
}

func SendRequest(req *http.Request) *http.Response {
	var resp *http.Response
	var err error

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	return resp
}

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
