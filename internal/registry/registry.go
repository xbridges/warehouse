package registry

import(
	"github.com/xbridges/warehouse/internal/domain/model/value"
	"github.com/xbridges/warehouse/internal/domain/repository"
	"github.com/xbridges/warehouse/internal/domain/service"

	"github.com/xbridges/warehouse/internal/infrastructure/datastore/adapter"

	"github.com/xbridges/warehouse/internal/application/usecase"

	"github.com/xbridges/warehouse/internal/interface/api/server/auth"
	"github.com/xbridges/warehouse/internal/interface/api/server/handler"
	"github.com/xbridges/warehouse/internal/interface/api/server/router"

)

type registry struct {
}

type Registry interface {
	// adapters
	NewKVSAdapter() adapter.KVSAdapter
	NewRDBAdapter() adapter.RDBAdapter

	// repositories
	NewSystemRepository(filepath string) (repository.SystemRepository, error)
	NewLoggingRepository(level string, destination string, lifecycle int) repository.LoggingRepository
	NewUserAuthRepository(adapter adapter.RDBAdapter) repository.UserAuthRepository
	NewCertificateRepository(adapter adapter.KVSAdapter) repository.CertificateRepository

	// services
	NewCertificateService(repo repository.CertificateRepository, authParam value.AuthParameter) service.CertificateService
	NewUserAuthService(repo repository.UserAuthRepository) service.UserAuthService

	// usecases
	NewLoginUsecase(usrserv service.UserAuthService, certserv service.CertificateService) usecase.LoginUsecase

	// handlers
	NewLoginHandler(usecase usecase.LoginUsecase) auth.LoginHandler

	// router
	NewRouter(handler handler.WarehouseHandler, cookieSercret string) router.Router
}

func NewRegistry() Registry{
	return &registry{}
}

func(r *registry) NewKVSAdapter() (adapter.KVSAdapter) {
	return adapter.NewKVSAdapter() 
}

func(r *registry) NewRDBAdapter() (adapter.RDBAdapter) {
	return adapter.NewRDBAdapter() 
}

func(r *registry) NewSystemRepository(filepath string) (repository.SystemRepository, error) {
	return repository.NewSystemRepository(filepath) 
}

func(r *registry) NewLoggingRepository(level string, destination string, lifecycle int) repository.LoggingRepository {
	return repository.NewLoggingRepository(level, destination, lifecycle)
}

func(r *registry) NewUserAuthRepository(adapter adapter.RDBAdapter) repository.UserAuthRepository {
	return repository.NewUserAuthRepository(adapter) 
}

func(r *registry) NewCertificateRepository(adapter adapter.KVSAdapter) repository.CertificateRepository {
	return repository.NewCertificateRepository(adapter)
}

func(r *registry) NewCertificateService(repo repository.CertificateRepository, authParam value.AuthParameter) service.CertificateService {
	return service.NewCertificateService(repo, authParam)
}

func(r *registry) NewUserAuthService(repo repository.UserAuthRepository) service.UserAuthService{
	return service.NewUserAuthService(repo)
}

func(r *registry) NewLoginUsecase(usrserv service.UserAuthService, certserv service.CertificateService) usecase.LoginUsecase{
	return usecase.NewLoginUsecase(usrserv, certserv)
}

func(r *registry) NewLoginHandler(usecase usecase.LoginUsecase) auth.LoginHandler {
	return auth.NewLoginHandler(usecase)
}

func(r *registry) NewRouter(handler handler.WarehouseHandler, cookieSercret string) router.Router {
	return router.NewRouter(handler, cookieSercret)
}
