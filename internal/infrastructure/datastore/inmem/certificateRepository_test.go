package inmem

import(
	"fmt"
	"time"
	"testing"
	"context"
	"github.com/xbridges/warehouse/internal/domain/model/value"
	"github.com/xbridges/warehouse/internal/infrastructure/datastore/adapter"
)

func TestCertificateRepository( t *testing.T){

	adp := adapter.NewKVSAdapter()
	// host string, passwd string, selector int) error
	err := adp.Open("172.22.10.110:6379", "", 1)
	if err != nil {
		t.Fatal(err)
	}
	repo := NewCertificateRepository(adp)

	ctx := context.TODO()
	id := 22
	expr := time.Now()
	recv := time.Now().Add(1)
	u := value.NewUser("test_1", 1, "uuid1", 1, &id, nil, nil, &expr, &recv )
	cert, err := repo.CreateCertificate(ctx, u, 30)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("token: %+v\n", cert)

	authu, err := repo.FindCertificateBy(ctx, cert.GetToken())
	if err != nil {
		t.Fatal(err)
	}
	json, err := authu.MarshalJson()
	if err != nil {
		t.Fatal()
	}
	fmt.Printf("user: %v\n", string(json))
}