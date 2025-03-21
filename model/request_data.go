package model

import "net/http"

type RequestData struct {
	Method  string
	Url     string
	Headers http.Header
	Body    string
}

func NewRequestData() *RequestData {
	return &RequestData{
		Method:  http.MethodGet,
		Headers: make(http.Header),
	}
}
