package persistence

import(
	"testing"
	"fmt"
)

func TestSystemRepository( t *testing.T ){

	sys, err := NewSystemRepository("../../../../etc/warehouse.yaml")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("RDBConnetctionString: %s\n", sys.System().RDBConnectionString())
	host, port, schema, user, pass := sys.System().KVSConnectionParams()
	fmt.Printf("KVS: %s, %d, %d, %v, %v\n", host, port, schema, *user, *pass)
	level, dest, cycle := sys.System().LogParams()
	fmt.Printf("Log: %s, %s, %d\n", level, dest, cycle)

}