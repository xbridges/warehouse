package persistence

import(
	"github.com/xbridges/warehouse/internal/infrastructure/filesystem/driver"
	"github.com/xbridges/warehouse/internal/infrastructure/filesystem/adapter"
)

type loggingRepository struct {
	adapter.LogAdapter
}

func NewLoggingRepository(level string, destination string, lifecycle int) *loggingRepository{
	return &loggingRepository{
		LogAdapter: func (lvl string, dest string, cycle int) adapter.LogAdapter {
			return driver.NewLogDriver(lvl, dest, cycle)
		}(level, destination, lifecycle),
	}
}

func SharedLogger() adapter.LogAdapter {
	return driver.SharedLogger()
}