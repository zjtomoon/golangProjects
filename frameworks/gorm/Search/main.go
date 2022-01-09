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

	//Where
	//原生sql

	//获取第一条匹配的记录 SELECT * FROM users WHERE name = 'jinzhu' limit 1;
	db.Where("name = ?", "jinzhu").First(&user)

	//获取所有匹配的记录
	db.Where("name = ?", "jinzhu").Find(&user)

	// <>
	db.Where("name <> ?", "jinzhu").Find(&user)

	// In
	db.Where("name in (?)", []string{"jinzhu", "jinzhu 2"}).Find(&user)

	//Like
	db.Where("name Like ?", "%jin%").Find(&user)

	//And
	db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&user)

	//Time
	db.Where("updated_at > ?", "lastweek").Find(&user)

	//Between
	db.Where("created_at Between ? AND ?", "lastweek", "today").Find(&user)

	//Struct & Map
	db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
	//// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 LIMIT 1;

	// Map
	db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&user)
	//// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

	//多主键 slice 查询
	db.Where([]int64{20, 21, 22}).Find(&user)
	//// SELECT * FROM users WHERE id IN (20, 21, 22);

	db.Where(&User{Name: "jinzhu", Age: 0}).Find(&user)
	//SELECT * FROM users WHERE name = "jinzhu";

	//Not
	db.Not("name", "jinzhu").First(&user)
	//// SELECT * FROM users WHERE name <> "jinzhu" LIMIT 1;

	//不包含
	db.Not("name", []string{"jinzhu", "jinzhu 2"}).Find(&user)
	//// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");

	//不在主键slice中
	db.Not([]int64{1, 2, 3}).First(&user)
	//// SELECT * FROM users WHERE id NOT IN (1,2,3);

	db.Not([]int64{}).First(&user)
	//// SELECT * FROM users;

	//原生sql
	db.Not("name = ?", "jinzhu").First(&user)
	//// SELECT * FROM users WHERE NOT(name = "jinzhu");

	//struct
	db.Not(User{Name: "jinzhu"}).First(&user)
	// SELECT * FROM users WHERE name <> "jinzhu";

	//Or

	db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&user)
	//// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';

	//Struct
	db.Where("name = 'jinzhu'").Or(User{Name: "jinzhu 2"}).Find(&user)
	//// SELECT * FROM users WHERE name = 'jinzhu' OR name = 'jinzhu 2';

	//Map
	db.Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "jinzhu 2"}).Find(&user)

	//行内条件查询

	//// 通过主键进行查询 (仅适用于主键是数字类型)
	db.First(&user, 23)
	//// SELECT * FROM users WHERE id = 23 LIMIT 1;

	//非数字类型的主键查询
	db.First(&user, "id = ?", "string_primary_key")
	//// SELECT * FROM users WHERE id = 'string_primary_key' LIMIT 1;

}
