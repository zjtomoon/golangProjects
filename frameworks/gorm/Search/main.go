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

	// 原生SQL
	db.Find(&user, "name = ?", "jinzhu")
	//// SELECT * FROM users WHERE name = "jinzhu";

	db.Find(&user, "name <> ? AND age > ?", "jinzhu", 20)
	//// SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;

	//struct
	db.Find(&user, User{Age: 20})
	//// SELECT * FROM users WHERE age = 20;

	//Map
	db.Find(&user, map[string]interface{}{"age": 20})
	//// SELECT * FROM users WHERE age = 20;

	//额外的查询选项
	//为查询 SQL 添加额外的选项
	db.Set("gorm:query_option", "FOR UPDATE").First(&user, 10)
	//// SELECT * FROM users WHERE id = 10 FOR UPDATE;

	//FirstOrInit
	//	获取第一条匹配的记录，或者通过给定的条件下初始一条新的记录（仅适用与于 struct 和 map 条件）。
	//为查询到
	db.FirstOrInit(&user, User{Name: "non_existing"})
	//// user -> User{Name: "non_existing"}

	//查询到
	db.Where(User{Name: "jinzhu"}).FirstOrInit(&user)
	//// user -> User{Id: 111, Name: "Jinzhu", Age: 20}

	db.FirstOrInit(&user, map[string]interface{}{"name": "jinzhu"})
	//// user -> User{Id: 111, Name: "Jinzhu", Age: 20}

	//Attrs
	//如果未找到记录，则使用参数初始化 struct
	//未查询到
	db.Where(User{Name: "non_existing"}).Attrs(User{Age: 20}).FirstOrInit(&user)
	//// SELECT * FROM USERS WHERE name = 'non_existing';
	//// user -> User{Name: "non_existing", Age: 20}

	db.Where(User{Name: "non_existing"}).Attrs("age", 20).FirstOrInit(&user)
	//// SELECT * FROM USERS WHERE name = 'non_existing';
	//// user -> User{Name: "non_existing", Age: 20}

	//查询到
	db.Where(User{Name: "jinzhu"}).Attrs(User{Age: 20}).FirstOrInit(&user)
	//// SELECT * FROM USERS WHERE name = jinzhu';
	//// user -> User{Id: 111, Name: "Jinzhu", Age: 20}

	//Asign
	//无论是否查询到数据，都将参数赋值给 struct
	//未查询到
	db.Where(User{Name: "non_existing"}).Assign(User{Age: 20}).FirstOrInit(&user)
	//// user -> User{Name: "non_existing", Age: 20}

	//查询到
	db.Where(User{Name: "jinzhu"}).Assign(User{Age: 30}).FirstOrInit(&user)
	//// SELECT * FROM USERS WHERE name = jinzhu';
	//// user -> User{Id: 111, Name: "Jinzhu", Age: 20}

	//FirstOrCreate
	//未查询到
	db.FirstOrCreate(&user, User{Name: "no_existing"})
	//// INSERT INTO "users" (name) VALUES ("non_existing");
	//// user -> User{Id: 112, Name: "non_existing"}

	//查询到
	db.Where(User{Name: "jinzhu"}).FirstOrCreate(&user)
	//// user -> User{Id: 111, Name: "Jinzhu"}

	//Attrs 如果未查询到记录，通过给定的参数赋值给 struct ，然后使用这些值添加一条记录。
	//未查询到
	db.Where(User{Name: "non_existing"}).Attrs(User{Age: 20}).FirstOrCreate(&user)
	//// SELECT * FROM users WHERE name = 'non_existing';
	//// INSERT INTO "users" (name, age) VALUES ("non_existing", 20);
	//// user -> User{Id: 112, Name: "non_existing", Age: 20}

	//查询到
	db.Where(User{Name: "jinzhu"}).Attrs(User{Age: 30}).FirstOrCreate(&user)
	//// SELECT * FROM users WHERE name = 'jinzhu';
	//// user -> User{Id: 111, Name: "jinzhu", Age:20}

	//Assign 无论是否查询到，都将其分配给记录，并保存到数据库中。
	//未查询到
	db.Where(User{Name: "non_existing"}).Assign(User{Age: 20}).FirstOrCreate(&user)
	//// SELECT * FROM users WHERE name = 'non_existing';
	//// INSERT INTO "users" (name, age) VALUES ("non_existing", 20);
	//// user -> User{Id: 112, Name: "non_existing", Age: 20}

	//查询到
	db.Where(User{Name: "jinzhu"}).Assign(User{Age: 30}).FirstOrCreate(&user)
	//// SELECT * FROM users WHERE name = 'jinzhu';
	//// UPDATE users SET age=30 WHERE id = 111;
	//// user -> User{Id: 111, Name: "jinzhu", Age: 30}

	//高级查询
	//子查询
	//db.Where("amount > ?",db.Table("orders").Select("AVG(amount)").Where("state = ?","paid").QueryExpr()).Find(&orders)
	// SELECT * FROM "orders"  WHERE "orders"."deleted_at" IS NULL AND (amount > (SELECT AVG(amount) FROM "orders"  WHERE (state = 'paid')));

	//查询 指定要从数据库检索的字段，默认情况下，将选择所有字段。
	db.Select("name,age").Find(&user)
	//// SELECT name, age FROM users;

	db.Select([]string{"name", "age"}).Find(&user)
	//// SELECT name, age FROM users;
	db.Table("users").Select("COALESCE(age,?)", 42).Rows()
	//// SELECT COALESCE(age,'42') FROM users;

	//Order 使用 Order 从数据库查询记录时，当第二个参数设置为 true 时，将会覆盖之前的定义条件。
	db.Order("age desc,name").Find(&user)
	//// SELECT * FROM users ORDER BY age desc, name;

	//多个排序条件
	db.Order("age desc").Order("name").Find(&user)
	//// SELECT * FROM users ORDER BY age desc, name;

	//重新排序
	//db.Order("age desc").Find(&user1).Order("age",true).Find(&users2)
	//// SELECT * FROM users ORDER BY age desc; (users1)
	//// SELECT * FROM users ORDER BY age; (users2)

	// Limit
	db.Limit(3).Find(&user)
	//// SELECT * FROM users LIMIT 3;

	// 用 -1 取消 LIMIT 限制条件
	//db.Limit(10).Find(&user1).Find(&user2)
	//// SELECT * FROM users LIMIT 10; (users1)
	//// SELECT * FROM users; (users2)

	//Offset
	//db.Offset(3).Find(&users)
	//// SELECT * FROM users OFFSET 3;

	// 用 -1 取消 OFFSET 限制条件
	//db.Offset(10).Find(&users1).Offset(-1).Find(&users2)
	//// SELECT * FROM users OFFSET 10; (users1)
	//// SELECT * FROM users; (users2)

	//Count
	//	db.Where("name = ?", "jinzhu").Or("name = ?", "jinzhu 2").Find(&users).Count(&count)
	//	//// SELECT * from USERS WHERE name = 'jinzhu' OR name = 'jinzhu 2'; (users)
	//	//// SELECT count(*) FROM users WHERE name = 'jinzhu' OR name = 'jinzhu 2'; (count)
	//
	//	db.Model(&User{}).Where("name = ?", "jinzhu").Count(&count)
	//	//// SELECT count(*) FROM users WHERE name = 'jinzhu'; (count)
	//
	//	db.Table("deleted_users").Count(&count)
	//	//// SELECT count(*) FROM deleted_users;

	//Group和Having
	//rows, err := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Rows()
	//for rows.Next() {
	//	...
	//}
	//
	//rows, err := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Rows()
	//for rows.Next() {
	//	...
	//}
	//
	//type Result struct {
	//	Date  time.Time
	//	Total int64
	//}
	//db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Scan(&results)

	//joins

	//rows, err := db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Rows()
	//for rows.Next() {
	//	...
	//}
	//
	//db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)
	//
	//// 多个关联查询
	//db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("credit_cards.number = ?", "411111111111").Find(&user)

	//pluck
	var ages []int64
	db.Find(&user).Pluck("age", &ages)

	var names []string
	db.Model(&User{}).Pluck("name", &names)

	db.Table("deleted_users").Pluck("name", &names)

	// Requesting more than one column? Do it like this:
	db.Select("name, age").Find(&user)

	//scan
	type Result struct {
		Name string
		Age  int
	}

	var result Result
	db.Table("users").Select("name, age").Where("name = ?", 3).Scan(&result)

	// Raw SQL
	db.Raw("SELECT name, age FROM users WHERE name = ?", 3).Scan(&result)

}
