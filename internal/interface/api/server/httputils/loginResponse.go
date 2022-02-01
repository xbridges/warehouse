package httputils

import(
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"

	"github.com/xbridges/warehouse/internal/domain/model"
)

const(
	CookieCertificateKey = "WarehouseSessionID"
)

type LoginResponse struct {
	response gin.H
}

func NewLoginResponse(status ResponseCode, session sessions.Session, cert model.Certificate, data string) HandlerResponse{

	// ここにcookieをセットする
	if cert != nil {
		session.Set(CookieCertificateKey, cert.GetToken())	
		session.Options(sessions.Options{
			Domain:     cert.GetDomain(),
			MaxAge:     cert.GetLifeCycle(), 
			Secure:     true,	// httpsのみ
			HttpOnly:   false,	// javascriptでもcookieを取得可に
		})
		session.Save()
	}
	return &LoginResponse{response: NewHandlerResponse(status, data)}
}

func(r *LoginResponse) Response() gin.H {
	return r.response
}
