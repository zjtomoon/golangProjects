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
	db.First(&user)

	user.Name = "jinzhu 2"
	user.Age = 100
	db.Save(&user)
	// UPDATE users SET name='jinzhu 2', age=100, birthday='2016-01-01', updated_at = '2013-11-17 21:34:10' WHERE id=111;

	//如果单个属性被更改了，更新它
	db.Model(&user).Update("name", "hello")
	//// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;

	//使用组合条件更新单个属性
	//db.Model(&user).Where("active = ?",true).Update("name","hello")
	//// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;

	// 使用 `map` 更新多个属性，只会更新那些被更改了的字段
	db.Model(&user).Updates(map[string]interface{}{"name": "hello1", "age": 18})

	// 使用 `struct` 更新多个属性，只会更新那些被修改了的和非空的字段
	db.Model(&user).Updates(User{Name: "hello2", Age: 18})
	//// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

	// 警告： 当使用结构体更新的时候, GORM 只会更新那些非空的字段
	// 例如下面的更新，没有东西会被更新，因为像 "", 0, false 是这些字段类型的空值
	db.Model(&user).Updates(User{Name: "", Age: 0})

	//更新选中的字段
	db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello3", "age": 19})
	//// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;

	db.Model(&user).Omit("name").Updates(map[string]interface{}{"name": "hello", "age": 20})
	//// UPDATE users SET age=18, actived=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

	//更新列钩子方法
	db.Table("users").Where("id IN (?)", []int{0, 11}).Updates(map[string]interface{}{"name": "hello", "age": 20})
	//// UPDATE users SET name='hello', age=18 WHERE id IN (10, 11);

	// 使用结构体更新将只适用于非零值，或者使用 map[string]interface{}
	db.Model(User{}).Updates(User{Name: "hello", Age: 18})
	//// UPDATE users SET name='hello', age=18;

	// 使用 `RowsAffected` 获取更新影响的记录数
	//db.Model(User{}).Updates(User{Name:"hello",Age:18}).RowsAffected

	//带有表达式的SQL更新
	//DB.Model(&product).Update("price", gorm.Expr("price * ? + ?", 2, 100))
	////// UPDATE "products" SET "price" = price * '2' + '100', "updated_at" = '2013-11-17 21:34:10' WHERE "id" = '2';
	//
	//DB.Model(&product).Updates(map[string]interface{}{"price": gorm.Expr("price * ? + ?", 2, 100)})
	////// UPDATE "products" SET "price" = price * '2' + '100', "updated_at" = '2013-11-17 21:34:10' WHERE "id" = '2';
	//
	//DB.Model(&product).UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
	////// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = '2';
	//
	//DB.Model(&product).Where("quantity > 1").UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
	////// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = '2' AND quantity > 1;

	// 在更新 SQL 语句中添加额外的 SQL 选项
	//db.Model(&user).Set("gorm:update_option", "OPTION (OPTIMIZE FOR UNKNOWN)").Update("name","hello")
	//// UPDATE users SET name='hello', updated_at = '2013-11-17 21:34:10' WHERE id=111 OPTION (OPTIMIZE FOR UNKNOWN);

}
