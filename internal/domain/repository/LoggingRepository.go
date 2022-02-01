package repository

import(
	"github.com/xbridges/warehouse/internal/infrastructure/filesystem/persistence"
	"github.com/xbridges/warehouse/internal/infrastructure/filesystem/adapter"
)

type LoggingRepository interface {
	adapter.LogAdapter
}

func NewLoggingRepository(level string, destination string, lifecycle int) LoggingRepository{
	return persistence.NewLoggingRepository(level, destination, lifecycle)
}

func SharedLogger() adapter.LogAdapter {
	return persistence.SharedLogger()
}