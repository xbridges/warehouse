module github.com/xbridges/warehouse/internal/infrastructure/datastore/persistence

go 1.14

require (
	github.com/xbridges/warehouse/internal/domain/model v0.0.1
	github.com/xbridges/warehouse/internal/domain/model/value v0.0.1
	github.com/xbridges/warehouse/internal/domain/model/Error v0.0.1
	github.com/xbridges/warehouse/internal/infrastructure/datastore/driver v0.0.1
	github.com/xbridges/warehouse/internal/infrastructure/datastore/adapter v0.0.1
	github.com/xbridges/warehouse/internal/infrastructure/datastore/entity v0.0.1
)

replace (
	github.com/xbridges/warehouse/internal/domain/model v0.0.1 => ../../../domain/model
	github.com/xbridges/warehouse/internal/domain/model/value v0.0.1 => ../../../domain/model/value
	github.com/xbridges/warehouse/internal/domain/model/Error v0.0.1 => ../../../domain/model/error
	github.com/xbridges/warehouse/internal/infrastructure/datastore/driver v0.0.1 => ../driver
	github.com/xbridges/warehouse/internal/infrastructure/datastore/adapter v0.0.1 => ../adapter
	github.com/xbridges/warehouse/internal/infrastructure/datastore/entity v0.0.1 => ../entity
)