package api

import "net/http"

type Request struct {
	Method string
	Url    string
}

func NewDefaultRequest() Request {
	return Request{
		Method: http.MethodGet,
	}
}
