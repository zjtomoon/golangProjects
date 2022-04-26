package models

import "log"

// 更新数据库密码
func UpdatePassword(name string, newpasswd string) {
	err := db.Table("user_table").Where("name = ?", name).Update("Password", newpasswd).Error
	if err != nil {
		log.Println("update failed,err:", err)
	}
}
