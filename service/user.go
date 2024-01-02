package service

import (
	"duoduo-gin/global"
	"duoduo-gin/model"
)

type User struct{}

func (u User) Find() (d []model.Users) {
	global.DUO_DB.Find(&d)
	return
}

func (u User) Create(userData model.Users) bool {
	res := global.DUO_DB.Create(&userData)
	if res.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}
