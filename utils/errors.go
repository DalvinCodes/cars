package utils

import "errors"

var (
	ErrCreatingObiect   = errors.New("error creating object")
	ErrDeletingObject   = errors.New("error deleting object")
	ErrRetrievingObject = errors.New("error retrieving object")
	ErrUpdatingObject   = errors.New("error updating object")
)

var (
	ErrNotFound   = errors.New("object not found")
	ErrEmptyInput = errors.New("empty input")
)
