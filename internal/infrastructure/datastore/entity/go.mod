modulemodule github.com/xbridges/warehouse/internal/infrastructure/datastore/entity

go 1.14

require (
	github.com/xbridges/warehouse/internal/domain/model v0.0.1
	github.com/xbridges/warehouse/internal/domain/model/value v0.0.1
	github.com/xbridges/warehouse/internal/domain/model/Error v0.0.1
)

replace (
	github.com/xbridges/warehouse/internal/domain/model v0.0.1 => ../../domain/model
	github.com/xbridges/warehouse/internal/domain/model/value v0.0.1 => ../../domain/model/value
	github.com/xbridges/warehouse/internal/domain/model/Error v0.0.1 => ../../domain/model/error
)

