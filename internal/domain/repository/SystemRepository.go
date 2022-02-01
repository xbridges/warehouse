package repository

import(
	"github.com/xbridges/warehouse/internal/domain/model"
	"github.com/xbridges/warehouse/internal/infrastructure/filesystem/persistence"
)

type SystemRepository interface {
	System() model.SystemParameter	// systemパラメータの取得
}

func NewSystemRepository(filepath string) (SystemRepository, error) {
	return persistence.NewSystemRepository(filepath)
}

