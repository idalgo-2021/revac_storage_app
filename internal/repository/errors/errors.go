package errors

import "errors"

var (
	ErrNotFound          = errors.New("not found")
	ErrInvalidInput      = errors.New("invalid input")
	ErrDatabaseOperation = errors.New("database operation failed")
)
