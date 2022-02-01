package value

import (
	"fmt"
)

type valueRDBParameter struct {
	Host		RDBHost
	Port		RDBPort
	Schema		RDBSchema
	User		RDBUser
	Password	RDBPassword
	SSLMode		RDBSSLMode
}

type RDBHost		string
type RDBPort		int
type RDBSchema		string
type RDBUser		string
type RDBPassword	string
type RDBSSLMode		string


type RDBParameter interface {
	ConnectionString() string
}

func NewRDBParameter(host string, port int, schema string, user string, password string, ssl string) RDBParameter {
	return &valueRDBParameter {
		Host:		RDBHost(host),
		Port:		RDBPort(port),
		Schema:		RDBSchema(schema),
		User:		RDBUser(user),
		Password:	RDBPassword(password),
		SSLMode:	RDBSSLMode(ssl),
	}
}

func (v *valueRDBParameter) ConnectionString() string {
	if v.Port < 0 || v.Port > 0xFFFF {
		v.Port = 5432
	}
	if v.SSLMode != "enable" {
		v.SSLMode = "disable"
	}
	return fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v sslmode=%v", v.User, v.Password, v.Host, v.Port, v.Schema, v.SSLMode)
}