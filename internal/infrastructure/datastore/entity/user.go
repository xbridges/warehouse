package entity

import(
	"encoding/json"
	"github.com/xbridges/warehouse/internal/domain/model/value"
	"github.com/xbridges/warehouse/internal/domain/model/Error"
)

const(
	UserDisabled int = iota
	UserEnabled
)

const(
	UserAdmin int = iota
	UserEditor
	UserViewer
	UserGuest
)

type User struct {
	ID				int				`json:"id"`
	DisplayID		string			`json:"display_id"`
	Name			string			`json:"name"`
	Permission		int 			`json:"permission"`
	UUID			string			`json:"uuid"`
	EmailAddress	string			`json:"email"`
	ExpireAt		sqlNullTime		`json:"expiration"`
	RecoveryAt		sqlNullTime		`json:"recovery"`
	Status			int				`json:"status"`
}

func NewUserEntity() *User {
	return &User{}
}

func (u *User) Unmarshal(data []byte) (error){
	if u != nil {
		err := json.Unmarshal(data, u)
		if err != nil {
			return Error.NewSystemFailedError("json.Unmarshal", err)
		}
		return nil
	}
	return Error.NewNoModelError("entity.user")
}

func (u *User) GetModel() value.User {
	expr := &u.ExpireAt.Time
	if u.ExpireAt.Valid == false {
		expr = nil
	}
	recv := &u.RecoveryAt.Time
	if u.RecoveryAt.Valid == false {
		recv = nil
	}
	return value.NewUser(u.DisplayID, u.Permission, u.UUID, u.Status, &u.ID, &u.Name, &u.EmailAddress, expr, recv)
}