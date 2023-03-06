package exceptions

import (
	"fmt"
	"net/http"
)

type BadRequest struct {
	Err     error
	Message string
}

func (e BadRequest) Unwrap() error {
	return e.Err
}

func (e BadRequest) StatusCode() int {
	return http.StatusBadRequest
}

func (e BadRequest) Error() string {
	if e.Err != nil {
		return "bad request: " + e.Err.Error()
	}
	return "bad request: " + e.Message
}

func (e BadRequest) ErrorWithFmt() string {
	if e.Err != nil {
		return fmt.Errorf("bad request: %w", e.Err).Error()
	}
	return "bad request: " + e.Message
}

func (e BadRequest) Msg() string {
	return e.Message
}

var ErrBadRequest BadRequest
