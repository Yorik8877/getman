package main

import "net/http"

type ResponseParts struct {
	StatusCode int                 `json:"status_code"` // HTTP status code
	Headers    map[string][]string `json:"headers"`     // Response headers
	Body       string              `json:"body"`        // Response body
}

func CutResponse(r *http.Response) *ResponseParts {
	var response ResponseParts
	response.StatusCode = r.StatusCode
	response.Headers = r.Header
	response.Body = ReadResponseBody(r)
	return &response
}
