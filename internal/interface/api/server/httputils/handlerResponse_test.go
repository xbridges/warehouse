package httputils

import(
	"fmt"
	"testing"
)

func TestHandlerResponse(t *testing.T){

	h1 := NewHandlerResponse(RespSuccess, "aaaa")
//	h2 := NewHandlerResponse(RespLoginFailed, "bbbb")
	h3 := NewHandlerResponse(RespUserMismatch, "cccc")
	h4 := NewHandlerResponse(RespPasswordMismatch, "dddd")
	h5 := NewHandlerResponse(RespUserExpired, "eeee")
	h6 := NewHandlerResponse(RespUserLocked, "ffff")
	h7 := NewHandlerResponse(RespPasswordExpired, "gggg")
	
	fmt.Printf("resp: %+v\n", h1)
//	fmt.Printf("resp: %+v\n", h2)
	fmt.Printf("resp: %+v\n", h3)
	fmt.Printf("resp: %+v\n", h4)
	fmt.Printf("resp: %+v\n", h5)
	fmt.Printf("resp: %+v\n", h6)
	fmt.Printf("resp: %+v\n", h7)
}