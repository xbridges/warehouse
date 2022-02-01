package Error

import(
)

type noAdapterError struct {
	ValueError
}

type NoAdapterError interface {
	ValueError
}

func NewNoAdapterError(adapterName string) error{
	return &noAdapterError{ValueError: NewError(NoAdapter, adapterName)}
}