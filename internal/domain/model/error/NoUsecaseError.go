package Error

import(
)

type noUsecaseError struct {
	ValueError
}

type NoUsecaseError interface {
	ValueError
}

func NewNoUsecaseError() error{
	return &noUsecaseError{ValueError: NewError(NoUsecase, "")}
}