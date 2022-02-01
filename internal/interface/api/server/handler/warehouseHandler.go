package handler

import(
	"github.com/xbridges/warehouse/internal/interface/api/server/auth"
)

type WarehouseHandler interface {
	auth.LoginHandler
}

