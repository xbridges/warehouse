package model

import(
	"github.com/xbridges/warehouse/internal/domain/model/value"
)

type datastore struct {
	RDBParameter	value.RDBParameter
	KVSParameter	value.KVSParameter
}

type systemParameter struct {
	Datastore	*datastore
	API			value.APIParameter
	Log			value.LogParameter
	Auth		value.AuthParameter
}

type SystemParameter interface {
	RDBConnectionString() string
	KVSConnectionParams() (string, int, int, *string, *string)
	LogParams() (string, string, int)
	APIParams() (int, int, string)
	AuthParam() value.AuthParameter
}

func NewSystemParameter(rdb value.RDBParameter, kvs value.KVSParameter, api value.APIParameter, log value.LogParameter, auth value.AuthParameter) SystemParameter {
	return &systemParameter{
		Datastore: &datastore{
			RDBParameter: rdb,
			KVSParameter: kvs,
		},
		API: api,
		Log: log,
		Auth: auth,
	}
}

func (s *systemParameter) RDBConnectionString() string {
	return s.Datastore.RDBParameter.ConnectionString()
}

func (s *systemParameter) KVSConnectionParams() (string, int, int, *string, *string) {
	return s.Datastore.KVSParameter.ConnectionParams()
}

func (s *systemParameter) LogParams() (string, string, int) {
	return s.Log.Params()
}

func (s *systemParameter) APIParams() (int, int, string) {
	return s.API.Params()
}

func (s *systemParameter) AuthParam() value.AuthParameter {
	return s.Auth
}
