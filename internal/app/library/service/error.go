package service

import "errors"

var (
	// ErrInvalidId ...
	ErrInvalidId = errors.New("invalid id")
)

// ErrInvalidId check is a ErrInvalidId
func IsErrInvalidId(err error) bool {
	return errors.Is(err, ErrInvalidId)
}
