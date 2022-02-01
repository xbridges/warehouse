package auth

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/xbridges/warehouse/internal/domain/model/Error"
	"github.com/xbridges/warehouse/internal/domain/model/value"
	"github.com/xbridges/warehouse/internal/application/usecase"
	"github.com/xbridges/warehouse/internal/interface/api/server/httputils"
)

type loginHandler struct {
	usecase usecase.LoginUsecase
	authParam value.AuthParameter
}

type LoginHandler interface {
	Login(c *gin.Context)
}

func NewLoginHandler(usecase usecase.LoginUsecase) LoginHandler{
	return &loginHandler{usecase: usecase}
}

func (h *loginHandler) Login(c *gin.Context) {
	if h == nil {
		resp := httputils.NewLoginResponse(httputils.RespSystemFailed, nil, nil, "handler is null")
		c.JSONP(http.StatusBadRequest, resp.Response())
		return
	}
	if h.usecase == nil {
		resp := httputils.NewLoginResponse(httputils.RespSystemFailed, nil, nil, "usecase is null")
		c.JSONP(http.StatusBadRequest, resp.Response())
		return	
	}
	lparam := httputils.LoginParam{}
	
	// インプット情報をバインド
	err := c.BindJSON(&lparam)
	if err != nil {
		resp := httputils.NewLoginResponse(httputils.RespSystemFailed, nil, nil, err.Error())
		c.JSONP(http.StatusBadRequest, resp.Response())
		return
	}

	// ログイン処理(TODO: user情報も返すこと)
	cert, user, err := h.usecase.Login(c.Request.Context(), lparam.ID, lparam.Password)	
	if err != nil {
		httpStatus := http.StatusBadRequest
		var resp httputils.HandlerResponse

		// errで応答を分岐する
		switch err.(type) {
		case Error.NotFoundError:
			switch err.(Error.NotFoundError).ErrorCode() {
			case Error.UserMismatch:
				resp = httputils.NewLoginResponse(httputils.RespUserMismatch, nil, nil, lparam.ID)
			case Error.PasswordMismatch:
				resp = httputils.NewLoginResponse(httputils.RespPasswordMismatch, nil, nil, "null")
			default:
				resp = httputils.NewLoginResponse(httputils.RespSystemFailed, nil, nil, err.Error())
			}
		case Error.UserExpiredError:
			resp = httputils.NewLoginResponse(httputils.RespUserExpired, nil, nil, "null")
		case Error.PasswordExpiredError:
			resp = httputils.NewLoginResponse(httputils.RespPasswordExpired, nil, nil, "null")
		default:
			httpStatus = http.StatusInternalServerError
			resp = httputils.NewLoginResponse(httputils.RespSystemFailed, nil, nil, err.Error())
		}
		c.JSONP(httpStatus, resp.Response())
		return
	}
	
	// cookieを設定してログインOKを返す。
	var resp httputils.HandlerResponse
	jstr, err := user.MarshalJson(value.MarshalSafe)
	if err != nil {
		resp = httputils.NewLoginResponse(httputils.RespSystemFailed, nil, nil, err.Error())
	} else {
		session := sessions.Default(c)
		resp = httputils.NewLoginResponse(httputils.RespSuccess, session, cert, string(jstr))
	}
	c.JSONP(http.StatusOK, resp.Response())
	return
}
