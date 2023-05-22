package api

import "net/http"

type Request struct {
	Method  string
	Url     string
	Headers http.Header
}

func NewDefaultRequest() Request {
	return Request{
		Method:  http.MethodGet,
		Headers: map[string][]string{},
	}
}
