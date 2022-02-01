package value

import (
)

type valueKVSParameter struct {
	Host		KVSHost
	Port		KVSPort
	Schema		KVSSchema
	User		KVSUser
	Password	KVSPassword
}

type KVSHost		string
type KVSPort		int
type KVSSchema		int
type KVSUser		*string
type KVSPassword	*string


type KVSParameter interface {
	ConnectionParams() (string, int, int, *string, *string)
}

func NewKVSParameter(host string, port int, schema int, user *string, password *string) KVSParameter {
	return &valueKVSParameter {
		Host:		KVSHost(host),
		Port:		KVSPort(port),
		Schema:		KVSSchema(schema),
		User:		KVSUser(user),
		Password:	KVSPassword(password),
	}
}

func (v *valueKVSParameter) ConnectionParams() (string, int, int, *string, *string) {
	return string(v.Host), int(v.Port), int(v.Schema), (*string)(v.User), (*string)(v.Password)
}