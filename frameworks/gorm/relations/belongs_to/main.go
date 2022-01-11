package belongs_to

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

//属于
//type User struct {
//	gorm.Model
//	Name string
//}
//
//// `Profile` 属于 `User`， `UserID` 是外键
//type Profile struct {
//	gorm.Model
//	UserID int
//	User   User
//	Name   string
//}

//外键
//type User struct {
//	gorm.Model
//	Name string
//}
//
//type Profile struct {
//	gorm.Model
//	Name      string
//	User      User `gorm:"foreignkey:UserRefer"` // 使用 UserRefer 作为外键
//	UserRefer string
//}

//关联外键

type User struct {
	gorm.Model
	Refer int
	Name  string
}

type Profile struct {
	gorm.Model
	Name      string
	User      User `gorm:"association_foreignkey:Refer"` // use Refer 作为关联外键
	UserRefer string
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

	user := User{Refer: 10, Name: "hello"}
	profile := Profile{Name: "profile", User: User{
		Refer: 11,
		Name:  "hello1",
	}, UserRefer: "11"}

	db.Model(&user).Related(&profile)
	//// SELECT * FROM profiles WHERE user_id = 111; // 111 is user's ID
}
