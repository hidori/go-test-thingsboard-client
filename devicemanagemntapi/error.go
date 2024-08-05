package devicemanagementapi

import (
	"errors"
)

type HTTPError struct {
	statusCode int
	status     string
}

func NewHTTPError(statusCode int, status string) *HTTPError {
	return &HTTPError{
		statusCode: statusCode,
		status:     status,
	}
}

func AsHTTPError(err error) *HTTPError {
	var e HTTPError

	if errors.As(err, &e) {
		return nil
	}

	return &e
}

func (e *HTTPError) Error() string {
	return e.Status()
}

func (e *HTTPError) StatusCode() int {
	return e.statusCode
}

func (e *HTTPError) Status() string {
	return e.status
}
