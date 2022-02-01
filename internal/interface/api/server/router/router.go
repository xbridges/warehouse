package router

import(
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/xbridges/warehouse/internal/interface/api/server/handler"
)

type router struct {
	Engine *gin.Engine
}

type Router interface {
	Run( portWithPrefix string )	
}

func NewRouter( handler handler.WarehouseHandler, cookie_sercret string) Router {
	
	engine := gin.Default()

	store := cookie.NewStore([]byte(cookie_sercret))
	engine.Use(sessions.Sessions("warehouse_apisession", store))

	// route作成
	engine.POST("/login", handler.Login)
	return &router{Engine: engine}
}

func (r *router) Run(portWithPrefix string) {
	r.Engine.Run(portWithPrefix)	
}