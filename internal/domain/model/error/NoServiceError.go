package Error

import(
)

type noServiceError struct {
	ValueError
}

type NoServiceError interface {
	ValueError
}

func NewNoServiceError() error{
	return &noServiceError{ValueError: NewError(NoService, "")}
}