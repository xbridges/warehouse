module github.com/xbridges/warehouse/internal/domain/model/mock

go 1.14

require (
	github.com/xbridges/warehouse/internal/domain/model/value v0.0.1
)

replace (
	github.com/xbridges/warehouse/internal/domain/model/value v0.0.1 => ../value
)