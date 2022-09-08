package business

import "errors"

var (
	ErrDataNotValidate error = errors.New("data not validate")
	ErrNotFound        error = errors.New("data not found")
	ErrConflict        error = errors.New("data already exists")
	ErrInvalidObjectID error = errors.New("invalid object id")
)

func ErrorNotValid(msg string) error {
	ErrDataNotValidate = errors.New(msg)

	return ErrDataNotValidate
}

func ErrorNotFound(msg string) error {
	ErrNotFound = errors.New(msg)

	return ErrNotFound
}

func ErrorConflict(msg string) error {
	ErrConflict = errors.New(msg)

	return ErrConflict
}
