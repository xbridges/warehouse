package Error

import(
	"time"
)

type passwordExpiredError struct {
	ValueError
}

type PasswordExpiredError interface {
	ValueError
}

func NewPasswordExpiredError(t time.Time) error{
	return &passwordExpiredError{ValueError: NewError(PasswordExpired, t)}
}

