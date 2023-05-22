package model

import "net/http"

type RequestData struct {
	Method  string
	Url     string
	Headers http.Header
	Body    string
}
