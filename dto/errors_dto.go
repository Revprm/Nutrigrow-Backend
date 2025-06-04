package dto

import "errors"

var (
	ErrCreate        = errors.New("failed to create record")
	ErrGetById       = errors.New("failed to get record by id")
	ErrNotFound      = errors.New("record not found")
	ErrUpdate        = errors.New("failed to update record")
	ErrDelete        = errors.New("failed to delete record")
	ErrInvalidData   = errors.New("invalid data")
	ErrAlreadyExists = errors.New("record already exists")
)
