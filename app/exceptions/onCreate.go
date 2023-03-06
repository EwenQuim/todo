package exceptions

// OnCreate
func OnCreate(err error) error {
	// var liteErr *sqlite.Error

	// if errors.As(err, &liteErr) && liteErr.Code() == 2067 {
	// 	return Conflict{Err: fmt.Errorf("error creating item: %w", err)}
	// }
	// if err != nil {
	// 	return fmt.Errorf("error creating item: %w", err)
	// }
	return nil
}
