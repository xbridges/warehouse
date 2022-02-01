module github.com/xbridges/warehouse/cmd/warehouseapi

go 1.14

require (
	github.com/jinzhu/copier v0.3.5 // indirect
	github.com/xbridges/warehouse/internal/application/usecase v0.0.1
	github.com/xbridges/warehouse/internal/domain/model v0.0.1
	github.com/xbridges/warehouse/internal/domain/model/Error v0.0.1
	github.com/xbridges/warehouse/internal/domain/model/mock v0.0.1
	github.com/xbridges/warehouse/internal/domain/model/value v0.0.1
	github.com/xbridges/warehouse/internal/domain/repository v0.0.1
	github.com/xbridges/warehouse/internal/domain/service v0.0.1
	github.com/xbridges/warehouse/internal/infrastructure/datastore/adapter v0.0.1
	github.com/xbridges/warehouse/internal/infrastructure/datastore/driver v0.0.1
	github.com/xbridges/warehouse/internal/infrastructure/datastore/entity v0.0.1
	github.com/xbridges/warehouse/internal/infrastructure/datastore/inmem v0.0.1
	github.com/xbridges/warehouse/internal/infrastructure/datastore/persistence v0.0.1
	github.com/xbridges/warehouse/internal/infrastructure/filesystem/adapter v0.0.1
	github.com/xbridges/warehouse/internal/infrastructure/filesystem/driver v0.0.1
	github.com/xbridges/warehouse/internal/infrastructure/filesystem/persistence v0.0.1
	github.com/xbridges/warehouse/internal/interface/api/server/auth v0.0.1
	github.com/xbridges/warehouse/internal/interface/api/server/handler v0.0.1
	github.com/xbridges/warehouse/internal/interface/api/server/httputils v0.0.1
	github.com/xbridges/warehouse/internal/interface/api/server/router v0.0.1
	github.com/xbridges/warehouse/internal/registry v0.0.1

)

replace (

	github.com/xbridges/warehouse/internal/application/usecase v0.0.1 => ../../internal/application/usecase
	github.com/xbridges/warehouse/internal/domain/model v0.0.1 => ../../internal/domain/model
	github.com/xbridges/warehouse/internal/domain/model/Error v0.0.1 => ../../internal/domain/model/Error
	github.com/xbridges/warehouse/internal/domain/model/mock v0.0.1 => ../../internal/domain/model/mock
	github.com/xbridges/warehouse/internal/domain/model/value v0.0.1 => ../../internal/domain/model/value
	github.com/xbridges/warehouse/internal/domain/repository v0.0.1 => ../../internal/domain/repository
	github.com/xbridges/warehouse/internal/domain/service v0.0.1 => ../../internal/domain/service
	github.com/xbridges/warehouse/internal/infrastructure/datastore/adapter v0.0.1 => ../../internal/infrastructure/datastore/adapter
	github.com/xbridges/warehouse/internal/infrastructure/datastore/driver v0.0.1 => ../../internal/infrastructure/datastore/driver
	github.com/xbridges/warehouse/internal/infrastructure/datastore/entity v0.0.1 => ../../internal/infrastructure/datastore/entity
	github.com/xbridges/warehouse/internal/infrastructure/datastore/inmem v0.0.1 => ../../internal/infrastructure/datastore/inmem
	github.com/xbridges/warehouse/internal/infrastructure/datastore/persistence v0.0.1 => ../../internal/infrastructure/datastore/persistence
	github.com/xbridges/warehouse/internal/infrastructure/filesystem/adapter v0.0.1 => ../../internal/infrastructure/filesystem/adapter

	github.com/xbridges/warehouse/internal/infrastructure/filesystem/driver v0.0.1 => ../../internal/infrastructure/filesystem/driver
	github.com/xbridges/warehouse/internal/infrastructure/filesystem/persistence v0.0.1 => ../../internal/infrastructure/filesystem/persistence

	github.com/xbridges/warehouse/internal/interface/api/server/auth v0.0.1 => ../../internal/interface/api/server/auth
	github.com/xbridges/warehouse/internal/interface/api/server/handler v0.0.1 => ../../internal/interface/api/server/handler
	github.com/xbridges/warehouse/internal/interface/api/server/httputils v0.0.1 => ../../internal/interface/api/server/httputils
	github.com/xbridges/warehouse/internal/interface/api/server/router v0.0.1 => ../../internal/interface/api/server/router

	github.com/xbridges/warehouse/internal/registry v0.0.1 => ../../internal/registry
)
