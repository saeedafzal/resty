package core

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/saeedafzal/resty/model"
)

func DoRequest(requestData *model.RequestData) (*model.ResponseData, error) {
	req, err := http.NewRequest(requestData.Method, requestData.Url, strings.NewReader(requestData.Body))
	if err != nil {
		return nil, err
	}

	req.Header = requestData.Headers

	// 10 second timeout for requests
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	start := time.Now()
	res, err := http.DefaultClient.Do(req.WithContext(ctx))
	elapsed := time.Since(start)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &model.ResponseData{
		StatusCode: res.StatusCode,
		Time:       elapsed.Milliseconds(),
		Body:       string(b),
		Headers:    res.Header,
	}, nil
}
