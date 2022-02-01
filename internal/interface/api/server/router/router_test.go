package router

import(
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/xbridges/warehouse/internal/interface/api/server/auth"
	"github.com/xbridges/warehouse/internal/interface/api/server/handler"
)

type loginhandlerMock struct {
}

func NewLoginhandlerMock() *loginhandlerMock{
	return &loginhandlerMock{}
}

func(mlogin *loginhandlerMock) Login(c *gin.Context){
	fmt.Println("Do something process.")
	return
}

func TestRouter(t *testing.T){

	router := NewRouter(
		func() handler.WarehouseHandler{
			return &struct {
				auth.LoginHandler
			}{
				LoginHandler: NewLoginhandlerMock(),
			}
		}(),
		"test_cookie",
	)

	router.Run(":8080")

}
