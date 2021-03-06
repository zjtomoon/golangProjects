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

//type User struct {
//	gorm.Model
//	Name string
//	Age  *int `gorm:"default:18"`
//}

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
	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	db.NewRecord(user) // => 返回 `true` ，因为主键为空
	db.Create(&user)
	db.NewRecord(user) // => 在 `user` 之后创建返回 `false`

	var animal = Animal{Age: 99, Name: ""}
	db.Create(&animal)
	// INSERT INTO animals("age") values('99');
	// SELECT name from animals WHERE ID=111; // 返回的主键是 111
	// animal.Name => 'galeone'

	//查询

	// 获取第一条记录，按主键排序 SELECT * FROM users ORDER BY id LIMIT 1;
	db.First(&user)
	//获取一条记录，不指定排序 SELECT * FROM users LIMIT 1;
	db.Take(&user)
	//获取最后一条记录，按主键排序  SELECT * FROM users ORDER BY id DESC LIMIT 1;
	db.Last(&user)
	//获取所有的记录 SELECT * FROM users;
	db.Find(&user)
	//通过主键进行查询 (仅适用于主键是数字类型)  SELECT * FROM users WHERE id = 10;
	db.First(&user, 10)
	db.Find(&user, 10)

}
