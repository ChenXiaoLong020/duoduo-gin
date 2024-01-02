package main

import (
	"duoduo-gin/global"
	"duoduo-gin/initialize"
)

func main() {

	global.DUO_CONFIG = initialize.ConfigInit()
	global.DUO_DB = initialize.GormInit()

	initialize.RouterInit()
}
