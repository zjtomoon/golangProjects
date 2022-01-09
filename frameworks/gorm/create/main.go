package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

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

	//创建
	db.DropTableIfExists(&User{})
	db.CreateTable(&User{})
	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	db.NewRecord(user) // => 返回 `true` ，因为主键为空
	db.Create(&user)
	db.NewRecord(user) // => 在 `user` 之后创建返回 `false`

	db.DropTableIfExists(&Animal{})
	db.CreateTable(&Animal{})
	var animal = Animal{Age: 99, Name: ""}
	db.Create(&animal)
	// INSERT INTO animals("age") values('99');
	// SELECT name from animals WHERE ID=111; // 返回的主键是 111
	// animal.Name => 'galeone'
}
