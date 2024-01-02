package middleware

import (
	"duoduo-gin/app/respone"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Error() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				respone.Fail(ctx, "系统错误")
				ctx.Abort()
				return
			}
		}()

		ctx.Next()

	}
}
