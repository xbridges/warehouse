module github.com/xbridges/warehouse/internal/domain/model/value

go 1.14

require (
	github.com/google/uuid v1.3.0
	github.com/jinzhu/copier v0.3.5
	github.com/xbridges/warehouse/internal/domain/model/Error v0.0.1
)

replace github.com/xbridges/warehouse/internal/domain/model/Error v0.0.1 => ../error
