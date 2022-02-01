package registry

import(
	"testing"
)

func TestRegistry(t *testing.T){

	r := NewRegistry()

	sys, err := r.NewSystemRepository("../../etc/warehouse.yaml")
	if err != nil {
		t.Fatal(err)
	}
	level, dest, cycle := sys.System().LogParams()
	log := r.NewLoggingRepository(level, dest, cycle)
	log.Info("test_registry", nil)
}