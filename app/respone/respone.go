package respone

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Respone struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Result(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Respone{code, msg, data})
}

func Success(c *gin.Context) {
	Result(c, 0, "操作成功", nil)
}

func Fail(c *gin.Context, msg string) {
	Result(c, 500, msg, nil)
}

func WithMessage(c *gin.Context, msg string) {
	Result(c, 0, msg, nil)
}

func WithData(c *gin.Context, data interface{}) {
	Result(c, 0, "操作成功", data)
}
