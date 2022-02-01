package repository

import (
	"fmt"
	"testing"
	"github.com/xbridges/warehouse/internal/infrastructure/datastore/adapter"

)

func TestRepository(t *testing.T){

	sys, err := NewSystemRepository("../../../etc/warehouse.yaml")
	if err != nil {
		t.Fatal(err)
	}
	
	fmt.Printf("RDBConnetctionString: %s\n", sys.System().RDBConnectionString())
	host, port, schema, user, pass := sys.System().KVSConnectionParams()
	fmt.Printf("KVS: %s, %d, %d, %v, %v\n", host, port, schema, *user, *pass)
	level, dest, cycle := sys.System().LogParams()
	fmt.Printf("Log: %s, %d\n", dest, cycle)

	syse, err := NewSystemRepository("warehouse.yaml")
	if err == nil {
		t.Fatal(err)
	}

	fmt.Printf("system parameter repository: %v\n", syse)

	l := NewLoggingRepository(level, dest, cycle)
	defer l.Sync()
	l.Info("Logging test!", map[string]interface{}{"count": 1})

	adapter := adapter.NewRDBAdapter()
	err = adapter.Open(sys.System().RDBConnectionString())
	if err != nil {
		t.Fatal(err)
	}
	defer adapter.Close()
	r := NewUserAuthRepository(adapter)
	u, err := r.FindUserBy("disp_kuro")
	if err != nil {
		t.Fatal(err)
	}
	p, tim, err := r.FindAuthorizationBy("4cf28004f60aa85952a5a0077f5c93037536dc6163c9a8f274146d1b0e654ba4")
	if err != nil {
		t.Fatal(err)
	}
	userJson, _ := u.Json()
	passphrase1, _ := u.GetPassphrase("1234")
	passphrase2, _ := u.GetPassphrase("5678")
	l.Info("userAuthRepo Test", map[string]interface{}{"user": string(userJson), "passphrase1": passphrase1, "passphrase2": passphrase2})
	l.Info("userAuthRepo Test", map[string]interface{}{"time": *tim, "passphrase": *p})	
}