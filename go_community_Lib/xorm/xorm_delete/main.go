package main

import (
	"fmt"
	"time"
	"xorm.io/xorm"
	_ "github.com/go-sql-driver/mysql"
)

type Player struct {
	Id int64
	Name string
	Age int
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func main() {
	engine,_ := xorm.NewEngine("mysql","root:123@/test?charset=utf8")

	engine.Sync2(&Player{})
	engine.Insert(&Player{Name: "dj",Age: 18})

	p := &Player{}
	engine.Where("name = ?","dj").Get(p)
	fmt.Println("after insert:",p)
	time.Sleep(5 * time.Second)

	engine.Table(&Player{}).ID(p.Id).Update(map[string]interface{}{"age":30})

	engine.Where("name = ?","dj").Get(p)
	fmt.Println("after update:",p)
	time.Sleep(5 * time.Second)

	engine.ID(p.Id).Delete(&Player{})

	engine.Where("name = ?","dj").Unscoped().Get(p)
	fmt.Println("after delete:",p)
}
