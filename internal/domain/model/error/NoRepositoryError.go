package Error

import(
)

type noRepositoryError struct {
	ValueError
}

type NoRepositoryError interface {
	ValueError
}

func NewNoRepositoryError() error{
	return &noRepositoryError{ValueError: NewError(NoRepository, "")}
}