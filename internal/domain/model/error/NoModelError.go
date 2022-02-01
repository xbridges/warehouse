package Error

import(
)

type noModelError struct {
	ValueError
}

type NoModelError interface {
	ValueError
}

func NewNoModelError(modelName string) error{
	return &noModelError{ValueError: NewError(NoModel, modelName)}
}