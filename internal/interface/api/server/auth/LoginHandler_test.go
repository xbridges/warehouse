package auth

import(
	"testing"

	"github.com/gin-gonic/gin"
	
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/xbridges/warehouse/internal/domain/model/value"
	"github.com/xbridges/warehouse/internal/domain/service"
	"github.com/xbridges/warehouse/internal/application/usecase"
	"github.com/xbridges/warehouse/internal/infrastructure/datastore/adapter"
	"github.com/xbridges/warehouse/internal/infrastructure/datastore/inmem"
	"github.com/xbridges/warehouse/internal/infrastructure/datastore/persistence"

//	"github.com/xbridges/warehouse/internal/interface/api/server/httputils"
)

func TestLoginHandler(t *testing.T) {

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
	
	authp := value.NewAuthParameter("/", "xbridges.com", 30, 1)

	certServ := service.NewCertificateService(inmem.NewCertificateRepository(kvs), authp)
	uauthServ := service.NewUserAuthService(persistence.NewUserAuthRepository(rdb))

	app := usecase.NewLoginUsecase(uauthServ, certServ)

	lhandler := NewLoginHandler(app)
	router := gin.Default()
	store := cookie.NewStore([]byte("sercret_test"))
	router.Use(sessions.Sessions("warehouse_apisession", store))

	router.POST("/login", lhandler.Login)
	router.Run(":8080")

	// goroutineで分けて接続するテストをしないといけないね。
}