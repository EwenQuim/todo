package exceptions_test

import (
	"errors"
	"testing"

	"github.com/EwenQuim/todo-app/app/exceptions"
)

var br = exceptions.BadRequest{Err: errors.New("error")}

func BenchmarkBadRequest_Error(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = br.Error()
	}
}

func BenchmarkBadRequest_ErrorWithFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = br.ErrorWithFmt()
	}
}
