package exceptions

import (
	"errors"
	"fmt"
	"net/http"
)

type NotFound struct {
	EntityInfo any
	Err        error
	EntityName string
	Message    string
}

func (e NotFound) Msg() string {
	return e.Message
}

func (e NotFound) Unwrap() error {
	return e.Err
}

func (e NotFound) StatusCode() int {
	return http.StatusNotFound
}

func (e NotFound) Error() string {
	if e.EntityName == "" {
		e.EntityName = "entity"
	}
	return fmt.Errorf("cannot find %s %v: %w", e.EntityName, e.EntityInfo, e.Err).Error()
}

var ErrNotFound = NotFound{Err: errors.New("not found")}
