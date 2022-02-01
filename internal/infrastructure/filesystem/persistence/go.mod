module github.com/xbridges/warehouse/internal/infrastructure/filesystem/persistence

go 1.14

require (
	github.com/goccy/go-yaml v1.9.4
	go.uber.org/zap v1.20.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0

	github.com/xbridges/warehouse/internal/domain/model v0.0.1
	github.com/xbridges/warehouse/internal/domain/model/value v0.0.1
	github.com/xbridges/warehouse/internal/infrastructure/filesystem/driver v0.0.1
	github.com/xbridges/warehouse/internal/infrastructure/filesystem/adapter v0.0.1
)

replace (
	github.com/xbridges/warehouse/internal/domain/model v0.0.1 => ../../../domain/model
	github.com/xbridges/warehouse/internal/domain/model/value v0.0.1 => ../../../domain/model/value
	github.com/xbridges/warehouse/internal/infrastructure/filesystem/driver v0.0.1 => ../driver
	github.com/xbridges/warehouse/internal/infrastructure/filesystem/adapter v0.0.1 => ../adapter
)
