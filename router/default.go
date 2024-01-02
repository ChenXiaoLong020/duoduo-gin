package router

import (
	"duoduo-gin/app/controller"

	"github.com/gin-gonic/gin"
)

func DefaultRouter(router *gin.Engine) {

	router.GET("/", controller.Home{}.Index)
	router.GET("/add", controller.Home{}.Add)

}
