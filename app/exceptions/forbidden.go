package exceptions

import (
	"fmt"
	"net/http"
)

type Forbidden struct {
	Err     error
	Message string
}

func (e Forbidden) Unwrap() error {
	return e.Err
}

func (e Forbidden) StatusCode() int {
	return http.StatusForbidden
}

func (e Forbidden) Error() string {
	if e.Err != nil {
		return fmt.Errorf("forbidden: %w", e.Err).Error()
	}
	return "forbidden: " + e.Message
}

func (e Forbidden) Msg() string {
	return e.Message
}

var ErrForbidden Forbidden
