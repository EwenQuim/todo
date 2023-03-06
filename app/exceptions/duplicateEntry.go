package exceptions

import (
	"fmt"
	"net/http"
)

type Conflict struct {
	Err error
}

func (e Conflict) Unwrap() error {
	return e.Err
}

func (e Conflict) StatusCode() int {
	return http.StatusConflict
}

func (e Conflict) Error() string {
	return fmt.Errorf("cannot create duplicate entry: %w", e.Err).Error()
}

var ErrConflict Conflict
