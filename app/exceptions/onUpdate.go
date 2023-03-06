package exceptions

import "fmt"

func OnUpdate(err error) error {
	return fmt.Errorf("cannot update: %w", err)
}
