package persistence

import(
	"fmt"
	"testing"
	"github.com/xbridges/warehouse/internal/infrastructure/datastore/adapter"
)

func TestAuthRepository(t *testing.T){
	
	a := adapter.NewRDBAdapter()
	err := a.Open("user=warehouse password=warehouse host=172.22.10.110 port=5432 dbname=warehousedb sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	r := NewUserAuthRepository(a)

	u1, err := r.FindUserBy("disp_kuro")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("user1: %+v\n", u1)

	u2, err := r.FindUserBy("disp_test")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("user2: %+v\n", u2)

	e1, err := r.FindUserBy("")
	if err == nil {
		t.Fatal()
	}
	fmt.Printf("user: %+v, err: %+v\n", e1, err)

	// 通常データ
	p1, t1, err := r.FindAuthorizationBy("3bb2e724dec4304ea770cb407df58617bfc54eb8c2264385dca1568d9046860c")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("pass1: %+v, time: %v\n", p1, t1)

	// 無期限
	p2, t2, err := r.FindAuthorizationBy("52990fa66fb09bd47ad547ad3e1d1346c6da74a541c0dec172ff529ce281c05c")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("pass2: %+v, time: %v\n", p2, t2)

	// 期限切れ
	p3, t3, err := r.FindAuthorizationBy("288b3b9aace34d77d203c28583ad9e2ec98bed96789c0cdf2923c25786551fed")
	if err == nil {
		t.Fatal(err)
	}
	fmt.Printf("pass3: %+v, time: %v, %+v\n", p3, t3, err)

}

