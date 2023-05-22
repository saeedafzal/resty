package api

import "net/http"

type Response struct {
	StatusCode   int
	ResponseTime int64
	Body         string
	Headers      http.Header
}
