package exceptions

import "fmt"

type SuperUserAlreadyExists struct {
	Err error
}

func (e SuperUserAlreadyExists) Unwrap() error {
	return e.Err
}

func (e SuperUserAlreadyExists) Error() string {
	return fmt.Errorf("super user already exists: %w", e.Err).Error()
}

var _ error = SuperUserAlreadyExists{}

var ErrSuperUserAlreadyExists SuperUserAlreadyExists
