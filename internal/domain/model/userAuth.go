package model

import(
	"github.com/xbridges/warehouse/internal/domain/model/value"
)

type userAuth struct {
	value.User
}

type UserAuth interface {
	value.User
}

func NewUserAuth(user value.User) UserAuth {
	return &userAuth{User: user}
}