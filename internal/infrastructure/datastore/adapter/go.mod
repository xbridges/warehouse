module github.com/xbridges/warehouse/internal/infrastructure/datastore/adapter

go 1.14

require (
	github.com/xbridges/warehouse/internal/infrastructure/datastore/driver v0.0.1
)

replace (
	github.com/xbridges/warehouse/internal/infrastructure/datastore/driver v0.0.1 => ../driver
)
