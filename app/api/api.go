package api

import (
	"io"
	"net/http"
	"time"

	"github.com/saeedafzal/resty/app/model"
)

type API struct {
	client http.Client
}

func NewAPI() API {
	return API{client: http.Client{}}
}

func (a API) DoRequest(requestModel model.RequestModel) (model.ResponseData, error) {
	responseData := model.ResponseData{}

	req, err := http.NewRequest(requestModel.Method, requestModel.Url, requestModel.Body)
	if err != nil {
		return responseData, err
	}

	headers := requestModel.Headers
	req.Header = headers
	headers.Set("User-Agent", "Resty/client-0.0.1")

	start := time.Now()
	res, err := a.client.Do(req)
	end := time.Now().Sub(start)

	if err != nil {
		return responseData, err
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return responseData, err
	}

	responseData.StatusCode = res.StatusCode
	responseData.Time = end.Milliseconds()
	responseData.Headers = res.Header
	responseData.Body = string(bytes)

	return responseData, nil
}
