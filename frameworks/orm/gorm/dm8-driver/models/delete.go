package models

import (
	"document_server/dao"
	"log"
)

func DropUserbyID(id int) {
	user := dao.User{}
	err := db.Table("user_table").Where("Id = ?", id).Delete(&user).Error
	if err != nil {
		log.Println("Failed to drop data,err:", err)
	}
}

func DropUserByName(name string) {
	user := dao.User{}
	err := db.Table("user_table").Where("UserName = ?", name).Delete(&user).Error
	if err != nil {
		log.Println("Failed to drop data,err:", err)
	}
}
