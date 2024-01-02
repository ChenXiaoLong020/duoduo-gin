package initialize

import (
	"duoduo-gin/global"
	"duoduo-gin/middleware"
	"duoduo-gin/router"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RouterInit() {

	//gin.SetMode(gin.ReleaseMode)

	app := gin.New()

	app.Use(middleware.Error())

	//指定静态文件目录
	app.Static("/static", "./static")

	router.DefaultRouter(app)

	app.Run(global.DUO_CONFIG.App.Host + ":" + strconv.Itoa(global.DUO_CONFIG.App.Port))
}
