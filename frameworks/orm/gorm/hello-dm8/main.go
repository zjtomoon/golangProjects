package main

import (
	"fmt"
	"github.com/encircles/gorm-dm8"
	_ "gitee.com/chunanyong/dm"
	"gorm.io/gorm"
)

func main() {

	// https://github.com/encircles/gorm-dm8
	dsn := "dm://SYSDBA:123456789@127.0.0.1:5236?ignoreCase=false&appName=wisdom&statEnable=false"
	_, err := gorm.Open(dm8.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("连接达梦数据库成功")

	//type Result struct {
	//	ID       string `gorm:"ID"`
	//	USERNAME string `gorm:"USERNAME"`
	//	PASSWORD string `gorm:"PASSWORD"`
	//}
	//
	//var result Result
	//
	//db.Raw("SELECT ID, USERNAME, PASSWORD FROM PERSON.PERSON WHERE ID = ?", "111").Scan(&result)
	//
	//fmt.Printf("%+v", result)
}
