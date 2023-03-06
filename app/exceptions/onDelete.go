package exceptions

import (
	"database/sql"
	"errors"
)

// OnDelete
func OnDelete(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return NotFound{Err: err}
	}
	return err
}
