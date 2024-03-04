package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int
	Msg  string
	Data interface{}
}

func Res(c *gin.Context, code int64, data interface{}, msg string) {
	httpCode := http.StatusOK
	if code != http.StatusOK {
		httpCode = int(code)
	}

	c.JSON(httpCode, Response{
		Code: int(code),
		Msg:  msg,
		Data: data,
	})
	return
}
