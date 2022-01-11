package Has_Many

import "github.com/jinzhu/gorm"

//一对多
//has many 关联就是创建和另一个模型的一对多关系， 不像 has one，所有者可以拥有0个或多个模型实例。
//
//例如，如果你的应用包含用户和信用卡， 并且每一个用户都拥有多张信用卡。

// 用户有多张信用卡，UserID 是外键
type User struct {
	gorm.Model
	CreditCards []CreditCard
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

//外键
//为了定义一对多关系， 外键是必须存在的，默认外键的名字是所有者类型的名字加上它的主键。
//
//就像上面的例子，为了定义一个属于User 的模型，外键就应该为 UserID。
//
//使用其他的字段名作为外键， 你可以通过 foreignkey 来定制它， 例如:
type User1 struct {
	gorm.Model
	CreditCards []CreditCard `gorm:"foreignkey:UserRefer"`
}

type CreditCard1 struct {
	gorm.Model
	Number    string
	UserRefer uint
}

//外键关联
//GORM 通常使用所有者的主键作为外键的值， 在上面的例子中，它就是 User 的 ID。
//
//当你分配信用卡给一个用户， GORM 将保存用户 ID 到信用卡表的 UserID 字段中。
//
//你能通过 association_foreignkey 来改变它， 例如:

type User2 struct {
	gorm.Model
	MemberNumber string
	CreditCards  []CreditCard `gorm:"foreignkey:UserMemberNumber;association_foreignkey:MemberNumber"`
}

type CreditCard2 struct {
	gorm.Model
	Number           string
	UserMemberNumber string
}

//多态关联
//支持多态的一对多和一对一关联。

type Cat struct {
	ID   int
	Name string
	Toy  []Toy `gorm:"polymorphic:Owner;"`
}

type Dog struct {
	ID   int
	Name string
	Toy  []Toy `gorm:"polymorphic:Owner;"`
}

type Toy struct {
	ID        int
	Name      string
	OwnerID   int
	OwnerType string
}

//使用一对多
//你可以通过Related 找到 has many 关联关系。

//db.Model(&user).Related(&emails)
//// SELECT * FROM emails WHERE user_id = 111; // 111 是用户表的主键
