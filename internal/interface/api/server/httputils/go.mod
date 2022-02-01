module github.com/xbridges/warehouse/internal/interface/api/server/httputils

go 1.14

require (
	github.com/gin-contrib/sessions v0.0.4
	github.com/gin-gonic/gin v1.7.7
	github.com/xbridges/warehouse/internal/domain/model v0.0.1
	github.com/xbridges/warehouse/internal/domain/model/value v0.0.1
	github.com/xbridges/warehouse/internal/domain/model/Error v0.0.1
)

replace (
	github.com/xbridges/warehouse/internal/domain/model v0.0.1 => ../../../../domain/model
	github.com/xbridges/warehouse/internal/domain/model/value v0.0.1 => ../../../../domain/model/value
	github.com/xbridges/warehouse/internal/domain/model/Error v0.0.1 => ../../../../domain/model/error
)
