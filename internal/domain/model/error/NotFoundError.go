package Error

import(
)

type notFoundError struct {
	ValueError
	err error
}

type NotFoundError interface {
	ValueError
	RuntimeError() error
}

func NewNotFoundError(code ErrorCode, src string, err error) error{
	switch code {
	case UserMismatch:
	case PasswordMismatch:
	case CertificateMismatch:
	default:
		// NotFoundError以外なシステムエラーとする。
		return &systemFailedError{ValueError: NewError(SystemFailed, src), err: err}
	}
	return &notFoundError{ValueError: NewError(code, src), err: err}
}

func(e *notFoundError) RuntimeError() error{
	return e.err
}
