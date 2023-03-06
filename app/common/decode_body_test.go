package common

import (
	"errors"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type testStruct struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,min=3,max=10"`
}

func (t testStruct) Normalize() (testStruct, error) {
	return t, nil
}

type notNormalizable struct{}

func (t notNormalizable) Normalize() (notNormalizable, error) {
	return t, errors.New("not normalizable")
}

func TestRequestBody(t *testing.T) {
	t.Run("RequestBody with request", func(t *testing.T) {
		body := `{"id": "test", "name": "test"}`
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/test", strings.NewReader(body))
		r.Header.Set("Content-Type", "application/json")
		_, err := RequestBody[testStruct](w, r)
		require.NoError(t, err)
	})

	t.Run("RequestBody without content-type", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/test", nil)
		_, err := RequestBody[testStruct](w, r)
		require.Error(t, err)
		require.ErrorAs(t, err, &HTTPError{})
		require.Equal(t, 415, err.(HTTPError).Code)
		require.Equal(t, "Content-Type header is not application/json", err.(HTTPError).Message)
	})

	t.Run("RequestBody with invalid content-type", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/test", nil)
		r.Header.Set("Content-Type", "text/plain")
		_, err := RequestBody[testStruct](w, r)
		require.Error(t, err)
		require.ErrorAs(t, err, &HTTPError{})
		require.Equal(t, 415, err.(HTTPError).Code)
		require.Equal(t, "Content-Type header is not application/json", err.(HTTPError).Message)
	})

	t.Run("RequestBody with invalid JSON body", func(t *testing.T) {
		body := `{ unclosed json`
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/test", strings.NewReader(body))
		r.Header.Set("Content-Type", "application/json")
		_, err := RequestBody[testStruct](w, r)
		require.Error(t, err)
	})

	t.Run("RequestBody with invalid request attributes", func(t *testing.T) {
		body := `{"id": "test", "name": "too long name"}`
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/test", strings.NewReader(body))
		r.Header.Set("Content-Type", "application/json")
		_, err := RequestBody[testStruct](w, r)
		require.Error(t, err)
	})

	t.Run("RequestBody with unnormalizable input", func(t *testing.T) {
		body := `{}`
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/test", strings.NewReader(body))
		r.Header.Set("Content-Type", "application/json")
		_, err := RequestBody[notNormalizable](w, r)
		require.Error(t, err)
	})
}

func FuzzRequestBody(f *testing.F) {
	f.Add("{}")

	f.Fuzz(func(t *testing.T, body string) {
		fmt.Println(body)
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/test", strings.NewReader(body))
		r.Header.Set("Content-Type", "application/json")

		RequestBody[testStruct](w, r)
	})
}
