package models

import (
	"dm8-driver/dao"
	"github.com/encircles/gorm-dm8"

	//"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strconv"
)

var db *gorm.DB

func init() {
	var err error
	//dns := "root:123456@tcp(127.0.0.1:3306)/document_server?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	//todo: 参考文档： https://eco.dameng.com/docs/zh-cn/app-dev/go-go.html

	dns := "dm://SYSDBA:123456789@127.0.0.1:5236?ignoreCase=false&appName=wisdom&statEnable=false"
	db, err = gorm.Open(dm8.Open(dns), &gorm.Config{})
	if err != nil {
		//panic("failed to connect to database")
		log.Fatalf("Failed to connect to database, err: %v", err)
	}

}

//todo: 测试
func FindUserName(id int) (name string) {
	var user dao.User
	err := db.Select("UserName").Where(dao.User{Id: id}).Find(&user).Error
	if err != nil {
		log.Println("Failed to find username,err:", err)
	}
	return user.UserName
}

// 根据id_username的格式创建表名
func CreateFileTableX(id int) {
	var name string
	var newName string
	name = FindUserName(id)
	newName = strconv.Itoa(id) + "_" + name
	err := db.Table(newName).AutoMigrate(&dao.FileTable{})
	if err != nil {
		return
	}
}

func CreateUserInfo(username string, password string) {
	user := dao.NewUser(username, password)
	err := db.Table("user_table").Create(user).Error
	if err != nil {
		log.Println("Failed to Insert userinfo", err)
	}
}
