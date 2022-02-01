package persistence

import(
	"fmt"
	"testing"
)

func TestLoggingRepository(t *testing.T){

	r := NewLoggingRepository("debug", "C:/Users/r-kuroda/Documents/02_repository/codeworkspace/go/warehouse/logwarehouse.log", 3)
	defer r.Sync()

	for i := 0; i<30000; i++{
		r.Info("Logging test!", map[string]interface{}{"count": i})
	}
	fmt.Printf("%+v\n", r)
}