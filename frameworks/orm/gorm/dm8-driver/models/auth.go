package models

import (
	"dm8-driver/dao"
	"log"
)

// todo: password直接保存为string是否安全？
func CheckAuth(name string, password string) bool {
	var user dao.User
	err := db.Select("password").Where("UserName = ?", name).Find(&user).Error
	if err != nil {
		log.Println("Failed to find,err:", err)
	}
	if user.Password == password {
		return true
	}
	return false
}
