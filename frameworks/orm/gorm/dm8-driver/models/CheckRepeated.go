package models

import (
	"document_server/dao"
	"log"
)

// 判断用户是否重复
func CheckRepeated(name string) bool {
	var user dao.User
	err := db.Table("user_table").Select("UserName").Where("UserName = ?", name).Find(&user).Error
	if err != nil {
		log.Println("Failed to find ,err: ", err)
	}
	if user.UserName == name {
		return true
	}
	return false
}
