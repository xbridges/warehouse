package Error

import(
)

type userExpiredError struct {
	ValueError
}

type UserExpiredError interface {
	ValueError
}

func NewUserExpiredError() error{
	return &userExpiredError{ValueError: NewError(UserExpired, "")}
}
