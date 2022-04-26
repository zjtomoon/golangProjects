package Has_One

import "github.com/jinzhu/gorm"

//Has one
//has one 关联也是与另一个模型建立一对一的连接，但语义（和结果）有些不同。 此关联表示模型的每个实例包含或拥有另一个模型的一个实例。
//
//例如，如果你的应用程序包含用户和信用卡，并且每个用户只能有一张信用卡。
// 用户有一个信用卡，CredtCardID 外键
type User struct {
	gorm.Model
	CreditCard CreditCard
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

//外键
//对于一对一关系，一个外键字段也必须存在，所有者将保存主键到模型关联的字段里。
//
//这个字段的名字通常由 belongs to model 的类型加上它的 primary key 产生的，就上面的例子而言，它就是 CreditCardID
//
//当你给用户一个信用卡， 它将保存一个信用卡的 ID 到 CreditCardID 字段中。
//
//如果你想使用另一个字段来保存这个关系，你可以通过使用标签 foreignkey 来改变它， 例如：
type User1 struct {
	gorm.Model
	CreditCard CreditCard `gorm:"foreignkey:CardRefer"`
}

type CreditCard1 struct {
	gorm.Model
	Number   string
	UserName string
}

//关联外键
//通常，所有者会保存 belogns to model 的主键到外键，你可以改为保存其他字段， 就像下面的例子使用 Number 。
type User2 struct {
	gorm.Model
	CreditCard CreditCard `gorm:"association_foreignkey:Number"`
}

type CreditCard2 struct {
	gorm.Model
	Number string
	UID    string
}

//多态关联
//支持多态的一对多和一对一关联。

type Cat struct {
	ID   int
	Name string
	Toy  Toy `gorm:"polymorphic:Owner;"`
}

type Dog struct {
	ID   int
	Name string
	Toy  Toy `gorm:"polymorphic:Owner;"`
}

type Toy struct {
	ID        int
	Name      string
	OwnerID   int
	OwnerType string
}

//使用一对一
//你可以通过 Related 找到 has one 关联。

//var card CreditCard
//db.Model(&user).Related(&card, "CreditCard")
////// SELECT * FROM credit_cards WHERE user_id = 123; // 123 是用户表的主键
//// CreditCard  是用户表的字段名，这意味着获取用户的信用卡关系并写入变量 card。
//// 像上面的例子，如果字段名和变量类型名一样，它就可以省略， 像：
//db.Model(&user).Related(&card)
