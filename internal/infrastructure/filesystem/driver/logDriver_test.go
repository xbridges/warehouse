package driver

import(
	"testing"
	//"fmt"
)

func TestLogRepository( t *testing.T ){

	SharedLogger().Info("test-shared logger is null.", nil)

	dest := "C:/Users/r-kuroda/Documents/02_repository/codeworkspace/go/warehouse/log/warehouse.log"
	cycle := 1
	repo := NewLogDriver("debug", dest, cycle)

	SharedLogger().Info("test-shared created.", nil)

	attribute := map[string]interface{}{"s1": "aaaa", "d1": 123}
	repo.Info("info", attribute)
	repo.Debug("debug", attribute)
	repo.Warn("warn", attribute)
	repo.Error("error", attribute)
	repo.Panic("panic", attribute)
//	repo.Fatal("fatal", attribute)

}