package repository

import "errors"

var (
	ErrRecipeNotFound = errors.New("recipe not found")
	ErrTaskNotFound   = errors.New("task not found")
	ErrInvalidUUID    = errors.New("invalid UUID")
)
