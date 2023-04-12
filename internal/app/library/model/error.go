package model

import (
	"errors"
)

var (
	// ErrInvalidString ...
	ErrInvalidString = errors.New("the provided string is not a valid ID")
	// ErrInvalidModel ...
	ErrInvalidModel = errors.New("invalid model")
)

// IsErrInvalidString check is a ErrInvalidString
func IsErrInvalidString(err error) bool {
	return errors.Is(err, ErrInvalidString)
}

// IsErrInvalidModel check is a ErrInvalidModel
func IsErrInvalidModel(err error) bool {
	return errors.Is(err, ErrInvalidModel)
}
