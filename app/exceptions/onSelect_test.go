package exceptions_test

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/EwenQuim/todo-app/app/exceptions"
	"github.com/stretchr/testify/require"
)

func TestOnSelect(t *testing.T) {
	t.Run("sql.ErrNoRows", func(t *testing.T) {
		err := sql.ErrNoRows
		require.Equal(t, exceptions.NotFound{Err: fmt.Errorf("error selecting item(s): %w", err)}, exceptions.OnSelect(err))
		require.ErrorAs(t, exceptions.OnSelect(err), &exceptions.ErrNotFound)
	})
	t.Run("idempotent with other errors", func(t *testing.T) {
		err := errors.New("other error")
		require.Equal(t, fmt.Errorf("error selecting item(s): %w", err), exceptions.OnSelect(err))
	})
	t.Run("nil", func(t *testing.T) {
		require.Nil(t, exceptions.OnSelect(nil))
	})
}
