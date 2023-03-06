package exceptions

import (
	"database/sql"
	"errors"
	"fmt"
)

// OnSelect
func OnSelect(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return NotFound{Err: fmt.Errorf("error selecting item(s): %w", err)}
	}
	if err != nil {
		return fmt.Errorf("error selecting item(s): %w", err)
	}
	return nil
}
