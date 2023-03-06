package exceptions

import (
	"fmt"
	"net/http"
)

type Unauthorized struct {
	Err     error
	Message string
}

func (e Unauthorized) Unwrap() error {
	return e.Err
}

func (e Unauthorized) StatusCode() int {
	return http.StatusUnauthorized
}

func (e Unauthorized) Error() string {
	return fmt.Errorf("unauthorized: %w", e.Err).Error()
}

func (e Unauthorized) Msg() string {
	return "unauthorized: " + e.Message
}

var ErrUnauthorized Unauthorized
