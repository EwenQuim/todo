package common

import (
	"errors"
	"net/http"
	"os"
)

type ErrorStatusCode interface {
	error
	StatusCode() int
}

type ErrorMsg interface {
	error
	Msg() string // A message that can be displayed to the end user
}

// HTTPError is an error that can be serialized to JSON.
type HTTPError struct {
	Err     error  `json:"-"`                 // The error that occurred, may be not privacy-friendly as it contains debug information
	Message string `json:"message,omitempty"` // a message to be displayed to the user
	Code    int    `json:"code"`
}

func (e HTTPError) StatusCode() int {
	if e.Code != 0 {
		return e.Code
	}

	var er ErrorStatusCode
	if errors.As(e.Err, &er) {
		return er.StatusCode()
	}

	return http.StatusInternalServerError
}

func (e HTTPError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	if e.Message != "" {
		return e.Message
	}
	if e.Code != 0 {
		return http.StatusText(e.Code)
	}
	return "internal server error"
}

func (e HTTPError) Msg() string {
	return e.Message
}

// SendError understands the error type and sends the appropriate response to the client.
func SendError(w http.ResponseWriter, err error) {
	e := HTTPError{Err: err}

	// Avoid nil dereferencing
	if err == nil {
		e.Err = errors.New("internal server error")
	}

	// If the error contains a status code, use it
	var errorWithStatusCode ErrorStatusCode
	if errors.As(e.Err, &errorWithStatusCode) {
		e.Code = errorWithStatusCode.StatusCode()
	}

	// Default to 500 if no status code is set
	if e.Code == 0 {
		e.Code = http.StatusInternalServerError
	}

	// If the error contains a message, use it
	var errorWithMessage ErrorMsg
	if errors.As(e.Err, &errorWithMessage) {
		e.Message = errorWithMessage.Msg()
	}

	// Default message relies on status code
	if e.Message == "" {
		if os.Getenv("ENV") == "dev" {
			e.Message = e.Error()
		} else {
			e.Message = http.StatusText(e.Code)
		}
	}

	SendJSON(w, e, e.Code)
}
