package httputils

import(
	"github.com/gin-gonic/gin"
)

type ResponseCode int
const (
	RespSuccess ResponseCode = iota
	RespSystemFailed
	RespUserMismatch
	RespPasswordMismatch
	RespUserExpired
	RespUserLocked
	RespPasswordExpired

)

var ResponseMessages map[ResponseCode]string = map[ResponseCode]string{
	RespSuccess: 			"success",
	RespSystemFailed:		"system failed",
	RespUserMismatch:		"login failed, id mismatch",
	RespPasswordMismatch:	"login failed, password mismatch",
	RespUserExpired:		"login failed, user expired",
	RespUserLocked: 		"login failed, user locked",
	RespPasswordExpired: 	"login failed, password expired",	
}

type handlerResponse struct{
}

type HandlerResponse interface{
	Response() gin.H
}

func NewHandlerResponse(status ResponseCode, data string) gin.H {
	return gin.H{
		"status": status,
		"message": ResponseMessages[status],
		"data": data,
	}
}