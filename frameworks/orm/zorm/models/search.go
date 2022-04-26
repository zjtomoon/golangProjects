package models

import (
	"fmt"
	"gitee.com/chunanyong/zorm"
	"testing"
)

// 1、查询单条数据到map
func TestQueryRowMap(t *testing.T) {
	//构造查询用的finder
	finder := zorm.NewSelectFinder("person") // select * from t_demo
	//finder.Append 第一个参数是语句,后面的参数是对应的值,值的顺序要正确.语句统一使用?,zorm会处理数据库的差异
	finder.Append("WHERE id=? and age in(?)", "3", []int{16, 17})
	//执行查询
	resultMap, err := zorm.QueryRowMap(ctx, finder)

	if err != nil { //标记测试失败
		t.Errorf("错误:%v", err)
	}
	//打印结果
	fmt.Println(resultMap)
}

// 2、查询单条到结构体
func TestQueryRow(t *testing.T) {

	type demoStruct struct {
		ID   int    `column:"id"`
		Name string `column:"name"`
		Age  int    `column:"age"`
	}

	//声明一个对象的指针,用于承载返回的数据
	demo := &demoStruct{}

	//构造查询用的finder
	finder := zorm.NewSelectFinder("person") // select * from t_demo
	//finder = zorm.NewSelectFinder(demoStructTableName, "id,userName") // select id,userName from t_demo
	//finder = zorm.NewFinder().Append("SELECT * FROM " + demoStructTableName) // select * from t_demo

	//finder.Append 第一个参数是语句,后面的参数是对应的值,值的顺序要正确.语句统一使用?,zorm会处理数据库的差异
	finder.Append("WHERE id=? ", "1")

	//执行查询
	_, err := zorm.QueryRow(ctx, finder, demo)

	if err != nil { //标记测试失败
		t.Errorf("错误:%v", err)
	}
	//打印结果
	fmt.Println(demo)
}

// 3、查询多条到map切片
func TestQueryMap(t *testing.T) {
	//构造查询用的finder
	finder := zorm.NewSelectFinder("person") // select * from person

	//创建分页对象,查询完成后,page对象可以直接给前端分页组件使用
	page := zorm.NewPage()
	page.PageNo = 1   //查询第1页,默认是1
	page.PageSize = 2 //每页20条,默认是20
	//执行查询
	listMap, err := zorm.QueryMap(ctx, finder, page)
	if err != nil { //标记测试失败
		t.Errorf("错误:%v", err)
	}
	//打印结果
	fmt.Println("总条数:", page.TotalCount, "  列表:", listMap)
}
