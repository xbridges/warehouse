package persistence

import (
	"io/ioutil"

	"github.com/xbridges/warehouse/internal/domain/model"
	"github.com/xbridges/warehouse/internal/domain/model/value"

	yaml "github.com/goccy/go-yaml"
)

type systemRepository struct {
	systemParameter model.SystemParameter
}

func NewSystemRepository( filepath string ) (*systemRepository, error) {

	systemConfig := struct{
		Database struct {
			Rdb struct {
				Host   		string	`yaml:"host"`
				Port   		int		`yaml:"port"`
				Schema 		string	`yaml:"schema"`
				User   		string 	`yaml:"user"`
				Password 	string 	`yaml:"password"`
				Ssl 		string	`yaml:"ssl"`
			} `yaml:"postgresql"`
			Kvs struct {
				Host   		string	`yaml:"host"`
				Port   		int		`yaml:"port"`
				Schema 		int		`yaml:"schema"`
				User   		*string `yaml:"user"`
				Password 	*string `yaml:"password"`
			} `yaml:"redis"`
		}	`yaml:"database"`
		Webapi struct {
			Version 	int 	`yaml:"version"`
			ListenPort 	int 	`yaml:"port"`
			SecretKey	string 	`yaml:"cookie_key"`
		}	`yaml:"webapi"`
		Log struct {
			LogLevel string     `yaml:"loglevel"`
			Destination string	`yaml:"destination"`
			LifeCycle 	int 	`yaml:"lifecycle"`
		}	`yaml:"log"`
		Authenticator struct {
			Domain 		string 	`yaml:"domain"`
			Expiration  int 	`yaml:"expiration"`
			Schema 		int 	`yaml:"schema"`
		} 	`yaml:"authenticator"`
		Application struct {
		}	`yaml:"Application"`
	}{}

	contents, err := ioutil.ReadFile( filepath )
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(contents, &systemConfig); err != nil {
		return nil, err	
	}

	config_values := model.NewSystemParameter(
		value.NewRDBParameter(
			systemConfig.Database.Rdb.Host,
			systemConfig.Database.Rdb.Port,
			systemConfig.Database.Rdb.Schema,
			systemConfig.Database.Rdb.User,
			systemConfig.Database.Rdb.Password,
			systemConfig.Database.Rdb.Ssl,
		),
		value.NewKVSParameter(
			systemConfig.Database.Kvs.Host,
			systemConfig.Database.Kvs.Port,
			systemConfig.Database.Kvs.Schema,
			systemConfig.Database.Kvs.User,
			systemConfig.Database.Kvs.Password,
		),
		value.NewAPIParameter(
			systemConfig.Webapi.Version,
			systemConfig.Webapi.ListenPort,
			systemConfig.Webapi.SecretKey,	
		),
		value.NewLogParameter(
			systemConfig.Log.LogLevel,
			systemConfig.Log.Destination,
			systemConfig.Log.LifeCycle,	
		),
		value.NewAuthParameter(
			systemConfig.Authenticator.Domain,
			systemConfig.Authenticator.Expiration,
			systemConfig.Authenticator.Schema,	
		),
	)

	rep := &systemRepository{ systemParameter: config_values }
	return rep, nil
}


func (r *systemRepository) System() model.SystemParameter {
	return r.systemParameter
}

