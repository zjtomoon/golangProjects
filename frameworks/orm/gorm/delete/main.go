package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

//GORM 默认使用 ID 作为主键名。
//  `gorm:"primary_key"`
type User struct {
	gorm.Model
	Name     string
	Age      int64
	Birthday time.Time
	Email    string `gorm:"type:varchar(100);unique_index"`
	Role     string `gorm:"size:255"`
	//MemberNumber *string `gorm:"unique;not null"`
	Num      int    `gorm:"AUTO_INCREMENT"`
	Address  string `gorm:"index:addr"` // 给Address 创建一个名字是  `addr`的索引
	IgnoreMe int    `gorm:"_"`          //忽略这个字段
}

type Animal struct {
	ID   int64
	Name string `gorm:"default:'galeone'"`
	Age  int64
}

func main() {
	db, err := gorm.Open("mysql", "tom:123@(127.0.0.1:3306)/gorm_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("--数据库连接失败")
	}
	fmt.Println("--连接数据库成功")
	defer db.Close()

	db.DropTableIfExists(&User{})
	db.CreateTable(&User{})

	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	//解除注释写入一条临时数据
	//db.Create(&user)

	////删除一条存在的记录
	db.Delete(&user)

	// 为删除 SQL 语句添加额外选项
	//db.Set("gorm:delete_option", "OPTION (OPTIMIZE FOR UNKNOWN)").Delete(&email)
	//// DELETE from emails where id=10 OPTION (OPTIMIZE FOR UNKNOWN);

	//批量删除
	//db.Where("email LIKE ?", "%jinzhu%").Delete(Email{})
	//// DELETE from emails where email LIKE "%jinzhu%";

	//db.Delete(Email{}, "email LIKE ?", "%jinzhu%")
	//// DELETE from emails where email LIKE "%jinzhu%";

	//软删除
	db.Delete(&user)
	//// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE id = 111;

	// 批量删除
	db.Where("age = ?", 20).Delete(&user)
	//// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE age = 20;

	//在查询记录时，软删除记录会被忽略
	db.Where("age = 20").Find(&user)
	//// SELECT * FROM users WHERE age = 20 AND deleted_at IS NULL;

	// 使用 Unscoped 方法查找软删除记录
	db.Unscoped().Where("age = 20").Find(&user)

	//使用unscoped方法永久删除记录
	//db.Unscoped().Delete(&order)
	//// DELETE FROM orders WHERE id=10;
}
