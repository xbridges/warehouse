package main

import(
	"fmt"
	"github.com/xbridges/warehouse/internal/registry"
	"github.com/xbridges/warehouse/internal/interface/api/server/auth"
	"github.com/xbridges/warehouse/internal/interface/api/server/handler"
)

func main() {
	
	r := registry.NewRegistry()

	sysrepo, err := r.NewSystemRepository("../../etc/warehouse.yaml")
	if err != nil {
		fmt.Printf("Application Error: %+v\n", err)
		return
	}

	level, logpath, cycle := sysrepo.System().LogParams()
	logrepo := r.NewLoggingRepository(level, logpath, cycle)
	logrepo.Info("START APPLICATION [warehouseapi]", nil)

	rdbAdapter := r.NewRDBAdapter()
	err = rdbAdapter.Open(sysrepo.System().RDBConnectionString())
	if err != nil {
		logrepo.Error("RDBConn failed.", map[string]interface{}{"ConnStr": sysrepo.System().RDBConnectionString()})
		return 	
	}
	defer rdbAdapter.Close()
	logrepo.Info("RDBConn succces", map[string]interface{}{"ConnStr": sysrepo.System().RDBConnectionString()})

	kvsAdapter := r.NewKVSAdapter()
	host, port, _, _, pass := sysrepo.System().KVSConnectionParams()	
	var argpass string
	if pass == nil {
		argpass = ""
	} else {
		argpass = *pass
	}
	
	// redis indexが異なる場合は、都度adapterを作成すること。
	authParam := sysrepo.System().AuthParam()
	err = kvsAdapter.Open(fmt.Sprintf("%s:%d", host, port), argpass, authParam.GetSchema())
	if err != nil {
		logrepo.Error("KVSConn failed.", map[string]interface{}{"host": host, "port": port, "schema": authParam.GetSchema()})
		return 	
	}
	logrepo.Info("KVSConn success.", map[string]interface{}{"host": host, "port": port, "schema": authParam.GetSchema()})
	
	_, port, secret := sysrepo.System().APIParams()
	r.NewRouter(
		func() handler.WarehouseHandler{
			return &struct {
				auth.LoginHandler
			}{
				r.NewLoginHandler(
					r.NewLoginUsecase(
						r.NewUserAuthService(r.NewUserAuthRepository(rdbAdapter)),
						r.NewCertificateService(r.NewCertificateRepository(kvsAdapter), sysrepo.System().AuthParam()),
					),
				),
			}
		}(),
		secret,
	).Run(fmt.Sprintf(":%d", port))
}
