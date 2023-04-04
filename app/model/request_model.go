package model

import (
	"io"
	"net/http"
)

type RequestModel struct {
	Method  string
	Url     string
	Headers http.Header
	Body    io.Reader
}

func NewRequestModel() RequestModel {
	return RequestModel{
		Method:  http.MethodGet,
		Url:     "",
		Headers: make(http.Header, 0),
	}
}
