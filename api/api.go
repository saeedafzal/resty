package api

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/saeedafzal/resty/model"
)

func DoRequest(requestData *model.RequestData) (*model.ResponseData, error) {
	// Body to buffer
	buffer := bytes.NewBuffer([]byte(requestData.Body))

	// Create request
	req, err := http.NewRequest(requestData.Method, requestData.Url, buffer)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	// Custom user-agent if not specified
	userAgent := requestData.Headers.Get("User-Agent")
	if userAgent == "" {
		requestData.Headers.Add("User-Agent", "RestyAgent/dev")
		defer requestData.Headers.Del("User-Agent")
	}

	// Set headers on the request
	req.Header = requestData.Headers

	// Do request
	start := time.Now()
	res, err := http.DefaultClient.Do(req)
	end := time.Now().Sub(start)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Read response body
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Return response
	return &model.ResponseData{
		StatusCode: res.StatusCode,
		Time:       end.Milliseconds(),
		Body:       string(b),
		Headers:    res.Header,
	}, nil
}
