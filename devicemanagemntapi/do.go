package devicemanagementapi

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

var clientDo = (&http.Client{}).Do

type empty struct{}

func Do[TRequest any, TResponse any](ctx context.Context, method string, url string, request *TRequest, statusCode int) (*TResponse, error) {
	var requestBody io.Reader

	if request != nil {
		b, err := json.Marshal(request)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		requestBody = bytes.NewReader(b)
	}

	req, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	res, err := clientDo(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != statusCode {
		return nil, errors.Errorf("Status: %s", res.Status)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var response TResponse

	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &response, nil
}

func Get[TResponse any](ctx context.Context, url string, statusCode int) (*TResponse, error) {
	response, err := Do[empty, TResponse](ctx, http.MethodGet, url, nil, statusCode)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return response, nil
}

func Post[TRequest any, TResponse any](ctx context.Context, url string, request *TRequest, statusCode int) (*TResponse, error) {
	response, err := Do[TRequest, TResponse](ctx, http.MethodPost, url, request, statusCode)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return response, nil
}

func Put[TRequest any, TResponse any](ctx context.Context, url string, request *TRequest, statusCode int) (*TResponse, error) {
	response, err := Do[TRequest, TResponse](ctx, http.MethodPut, url, request, statusCode)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return response, nil
}

func Delete[TResponse any](ctx context.Context, url string, statusCode int) (*TResponse, error) {
	response, err := Do[empty, TResponse](ctx, http.MethodDelete, url, nil, statusCode)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return response, nil
}
