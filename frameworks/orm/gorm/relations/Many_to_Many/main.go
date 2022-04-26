package Many_to_Many

import "github.com/jinzhu/gorm"

//多对多

//多对多为两个模型增加了一个中间表。
//
//例如，如果你的应用包含用户和语言， 一个用户会说多种语言，并且很多用户会说一种特定的语言。

// 用户拥有并属于多种语言， 使用  `user_languages` 作为中间表
type User struct {
	gorm.Model
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name string
}

//反向关联

// 用户拥有并且属于多种语言，使用 `user_languages` 作为中间表
type User1 struct {
	gorm.Model
	Languages []*Language `gorm:"many2many:user_languages;"`
}

type Language1 struct {
	gorm.Model
	Name  string
	Users []*User `gorm:"many2many:user_languages;"`
}

//db.Model(&language).Related(&users)
//// SELECT * FROM "users" INNER JOIN "user_languages" ON "user_languages"."user_id" = "users"."id" WHERE  ("user_languages"."language_id" IN ('111'))

//外键
type CustomizePerson struct {
	IdPerson string             `gorm:"primary_key:true"`
	Accounts []CustomizeAccount `gorm:"many2many:PersonAccount;association_foreignkey:idAccount;foreignkey:idPerson"`
}

type CustomizeAccount struct {
	IdAccount string `gorm:"primary_key:true"`
	Name      string
}

//中间表外键
//如果你想改变中间表的外键，你可以用标签 association_jointable_foreignkey, jointable_foreignkey

type CustomizePerson1 struct {
	IdPerson string             `gorm:"primary_key:true"`
	Accounts []CustomizeAccount `gorm:"many2many:PersonAccount;foreignkey:idPerson;association_foreignkey:idAccount;association_jointable_foreignkey:account_id;jointable_foreignkey:person_id;"`
}

type CustomizeAccount1 struct {
	IdAccount string `gorm:"primary_key:true"`
	Name      string
}

// 自引用
//为了定义一个自引用的多对多关系，你不得不改变中间表的关联外键。
//
//和来源表外键不同的是它是通过结构体的名字和主键生成的，例如：

type User2 struct {
	gorm.Model
	Friends []*User `gorm:"many2many:friendships;association_jointable_foreignkey:friend_id"`
}

//GORM 将创建一个带外键 user_id 和 friend_id 的中间表， 并且使用它去保存用户表的自引用关系。
//
//然后你可以像普通关系一样操作它， 例如：

//DB.Preload("Friends").First(&user, "id = ?", 1)
//
//DB.Model(&user).Association("Friends").Append(&User{Name: "friend1"}, &User{Name: "friend2"})
//
//DB.Model(&user).Association("Friends").Delete(&User{Name: "friend2"})
//
//DB.Model(&user).Association("Friends").Replace(&User{Name: "new friend"})
//
//DB.Model(&user).Association("Friends").Clear()
//
//DB.Model(&user).Association("Friends").Count()

//使用多对多
//db.Model(&user).Related(&languages, "Languages")
////// SELECT * FROM "languages" INNER JOIN "user_languages" ON "user_languages"."language_id" = "languages"."id" WHERE "user_languages"."user_id" = 111
//
////  当查询用户时预加载 Language
//db.Preload("Languages").First(&user)
