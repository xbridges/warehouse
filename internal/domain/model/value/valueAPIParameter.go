package value

import (
)

type valueAPIParameter struct {
	Version 	APIVersion
	ListenPort 	APIPort
	SecretKey	APISecretKey
}

type APIVersion		int
type APIPort		int
type APISecretKey	string

type APIParameter interface {
	Params() (int, int, string)
}

func NewAPIParameter(version int, port int, api_key string) APIParameter {
	if port > 0xFFFF || port < 0 {
		port = 58080
	}
	return &valueAPIParameter {
		Version:	APIVersion(version),
		ListenPort:	APIPort(port),
		SecretKey:	APISecretKey(api_key),
	}
}

func(v *valueAPIParameter) Params() (int, int, string){
	return int(v.Version), int(v.ListenPort), string(v.SecretKey)
}