package service

import(
	"fmt"
	"context"
	"testing"
	"github.com/xbridges/warehouse/internal/domain/model"
	"github.com/xbridges/warehouse/internal/domain/model/value"
	"github.com/xbridges/warehouse/internal/infrastructure/datastore/adapter"
	"github.com/xbridges/warehouse/internal/infrastructure/datastore/inmem"

)

func TestCertificateService(t *testing.T){

	a := adapter.NewKVSAdapter()
	err := a.Open("172.22.10.110:6379", "", 1)
	if err != nil {
		t.Fatal(err)
	}

	authp := value.NewAuthParameter("/", "domain", 30, 1)

	ctx := context.TODO()
	s := NewCertificateService(inmem.NewCertificateRepository(a), authp)

	authUser := model.NewUserAuth(value.NewUser("test_1", 1, "uuid1", 1, nil, nil, nil, nil, nil))
	cert, err := s.IssueCertificate(ctx, authUser)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("cert: %v, %v\n", cert.GetToken(), cert.GetExpiration())

}