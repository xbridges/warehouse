module github.com/xbridges/warehouse/internal/domain/repository

go 1.14

require (
	github.com/xbridges/warehouse/internal/domain/model v0.0.1
	github.com/xbridges/warehouse/internal/domain/model/value v0.0.1
	github.com/xbridges/warehouse/internal/infrastructure/filesystem/driver v0.0.1
	github.com/xbridges/warehouse/internal/infrastructure/filesystem/adapter v0.0.1
	github.com/xbridges/warehouse/internal/infrastructure/filesystem/persistence v0.0.1
	github.com/xbridges/warehouse/internal/infrastructure/datastore/driver v0.0.1
	github.com/xbridges/warehouse/internal/infrastructure/datastore/adapter v0.0.1
	github.com/xbridges/warehouse/internal/infrastructure/datastore/persistence v0.0.1
	github.com/xbridges/warehouse/internal/infrastructure/datastore/entity v0.0.1
	go.uber.org/zap v1.20.0
)

replace (
	github.com/xbridges/warehouse/internal/domain/model v0.0.1 => ../model
	github.com/xbridges/warehouse/internal/domain/model/value v0.0.1 => ../model/value
	github.com/xbridges/warehouse/internal/infrastructure/filesystem/driver v0.0.1 => ../../infrastructure/filesystem/driver
	github.com/xbridges/warehouse/internal/infrastructure/filesystem/adapter v0.0.1 => ../../infrastructure/filesystem/adapter
	github.com/xbridges/warehouse/internal/infrastructure/filesystem/persistence v0.0.1 => ../../infrastructure/filesystem/persistence
	github.com/xbridges/warehouse/internal/infrastructure/datastore/driver v0.0.1 => ../../infrastructure/datastore/driver
	github.com/xbridges/warehouse/internal/infrastructure/datastore/adapter v0.0.1 => ../../infrastructure/datastore/adapter
	github.com/xbridges/warehouse/internal/infrastructure/datastore/persistence v0.0.1 => ../../infrastructure/datastore/persistence
	github.com/xbridges/warehouse/internal/infrastructure/datastore/entity v0.0.1 => ../../infrastructure/datastore/entity
)
