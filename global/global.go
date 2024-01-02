package global

import (
	"duoduo-gin/config"

	"gorm.io/gorm"
)

var (
	//配置文件信息
	DUO_CONFIG config.Server
	DUO_DB     *gorm.DB
)
