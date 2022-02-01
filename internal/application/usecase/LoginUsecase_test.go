package usecase

import(
	"fmt"
	"context"
	"testing"

	"github.com/xbridges/warehouse/internal/domain/model/value"
	"github.com/xbridges/warehouse/internal/domain/service"
	"github.com/xbridges/warehouse/internal/infrastructure/datastore/adapter"
	"github.com/xbridges/warehouse/internal/infrastructure/datastore/inmem"
	"github.com/xbridges/warehouse/internal/infrastructure/datastore/persistence"
)


func TestLoginUsecase(t *testing.T){

	kvs := adapter.NewKVSAdapter()
	err := kvs.Open("172.22.10.110:6379", "", 1)
	if err != nil {
		t.Fatal(err)
	}

	rdb := adapter.NewRDBAdapter()
	err = rdb.Open("user=warehouse password=warehouse host=172.22.10.110 port=5432 dbname=warehousedb sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer rdb.Close()
	
	authp := value.NewAuthParameter("/", "domain", 30, 1)

	ctx := context.TODO()
	certServ := service.NewCertificateService(inmem.NewCertificateRepository(kvs), authp)
	uauthServ := service.NewUserAuthService(persistence.NewUserAuthRepository(rdb))

	app := NewLoginUsecase(uauthServ, certServ)
	cert, err := app.Login(ctx, "disp_kuro", "test")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("certificate: %s\n", *cert)

}