package service

import(
	"fmt"
	"testing"
	"github.com/xbridges/warehouse/internal/infrastructure/datastore/adapter"
	"github.com/xbridges/warehouse/internal/infrastructure/datastore/persistence"

)

func TestUserAuthService(t *testing.T){

	a := adapter.NewRDBAdapter()
	err := a.Open("user=warehouse password=warehouse host=172.22.10.110 port=5432 dbname=warehousedb sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	s := NewUserAuthService(persistence.NewUserAuthRepository(a))

	u, err := s.ConfirmUser("disp_kuro", "test")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", u)
	jstr, err := u.MarshalJson()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("json: %s\n", jstr)

}