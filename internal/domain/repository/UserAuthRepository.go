package repository

import(
	"time"
	"github.com/xbridges/warehouse/internal/domain/model"
	"github.com/xbridges/warehouse/internal/infrastructure/datastore/adapter"
	"github.com/xbridges/warehouse/internal/infrastructure/datastore/persistence"
)

type UserAuthRepository interface {
	FindUserBy(showid string) (model.UserAuth, error)
	FindAuthorizationBy(passphraes string) (*string, *time.Time, error)
}

func NewUserAuthRepository(adapter adapter.RDBAdapter) UserAuthRepository {
	return persistence.NewUserAuthRepository(adapter)
}
