package model

import "net/http"

type ResponseData struct {
	StatusCode int
	Time       int64
	Body       string
	Headers    http.Header
}
