package api

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/saeedafzal/resty/model"
)

type API struct {
	client http.Client
}

func NewAPI() API {
	return API{http.Client{}}
}

func (a API) DoRequest(requestData model.RequestData) (model.ResponseData, error) {
	buffer := bytes.NewBuffer([]byte(requestData.Body))

	req, err := http.NewRequest(requestData.Method, requestData.Url, buffer)
	defer a.closeResources(req.Body)
	if err != nil {
		return model.ResponseData{}, err
	}

	// Custom user agent (if not specified)
	userAgent := requestData.Headers.Get("User-Agent")
	if userAgent == "" {
		requestData.Headers.Add("User-Agent", "RestyAgent/dev")
		defer requestData.Headers.Del("User-Agent")
	}

	req.Header = requestData.Headers

	start := time.Now()
	res, err := a.client.Do(req)
	if err != nil {
		return model.ResponseData{}, err
	}
	end := time.Now().Sub(start)
	defer a.closeResources(res.Body)

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return model.ResponseData{}, err
	}

	return model.ResponseData{
		StatusCode: res.StatusCode,
		Time:       end.Milliseconds(),
		Body:       string(b),
		Headers:    res.Header,
	}, nil
}

func (a API) closeResources(buffer io.ReadCloser) {
	_ = buffer.Close()
}
