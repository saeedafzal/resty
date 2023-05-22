package api

import (
	"io"
	"net/http"
	"time"
)

type Api struct {
	client http.Client
}

func NewApi() Api {
	return Api{client: http.Client{}}
}

func (a Api) DoRequest(request Request) (Response, error) {
	var response Response

	// NOTE: Setup request
	req, err := http.NewRequest(request.Method, request.Url, nil)
	if err != nil {
		return response, err
	}

	// NOTE: Execute request
	start := time.Now()
	res, err := a.client.Do(req)
	end := time.Now().Sub(start)
	if err != nil {
		return response, err
	}

	// NOTE: Build and return response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return response, err
	}
	response.Body = string(body)

	response.StatusCode = res.StatusCode
	response.ResponseTime = end.Milliseconds()
	response.Headers = res.Header

	return response, nil
}
