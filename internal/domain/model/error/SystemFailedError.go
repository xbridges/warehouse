package Error

import(
)

type systemFailedError struct {
	ValueError
	err error
}

type SystemFailedError interface {
	ValueError
	RuntimeError() error
}

func NewSystemFailedError(src string, err error) error{
	return &systemFailedError{ValueError: NewError(SystemFailed, src), err: err}
}

func(e *systemFailedError) RuntimeError() error{
	return e.err
}
