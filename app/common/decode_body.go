package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/EwenQuim/todo-app/app/exceptions"
)

const (
	maxBodySize = 1048576
)

// Body deserializes the request body into the given type.
// If the request body is empty, it returns an error and write it to the response.
func RequestBody[T any](w http.ResponseWriter, r *http.Request) (T, error) {
	var t T

	// Deserialize the request body
	if r.Header.Get("Content-Type") != "application/json" {
		err := HTTPError{
			Code:    http.StatusUnsupportedMediaType,
			Message: "Content-Type header is not application/json",
			Err:     errors.New("Content-Type header is not application/json"),
		}
		SendError(w, err)
		return t, err
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxBodySize)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&t)
	if err != nil {
		errWrapped := fmt.Errorf("cannot decode request body to %T: %w", t, err)
		SendError(w, exceptions.BadRequest{Err: errWrapped, Message: errWrapped.Error()})
		return t, errWrapped
	}

	// Validate input
	// err = validation.Validate(validator, t)
	// if err != nil {
	// 	errWrapped := fmt.Errorf("cannot validate request body: %w", err)
	// 	common.SendError(w, exceptions.BadRequest{Err: err, Message: err.Error()})
	// 	return t, errWrapped
	// }

	return t, nil
}
