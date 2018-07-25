package kvs

import "errors"

var (
	ErrNotFound          = errors.New("not found")
	ErrInvalidType       = errors.New("invalid type")
	ErrInvalidComparison = errors.New("invalid comparison")
)
