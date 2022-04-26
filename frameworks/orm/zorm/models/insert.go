package models

import (
	"context"
	"gitee.com/chunanyong/zorm"
	"testing"
)

//定义结构体
type demoStruct struct {
	zorm.EntityStruct        //实现zrom.EntityStruct的两个方法
	ID                int    `column:"id"`
	Name              string `column:"name"`
	Age               int    `column:"age"`
}

//GetTableName 获取表名称
//IEntityStruct 接口的方法,实体类需要实现!!!
func (entity *demoStruct) GetTableName() string {
	return "person"
}

//GetPKColumnName 获取数据库表的主键字段名称.因为要兼容Map,只能是数据库的字段名称
//不支持联合主键,变通认为无主键,业务控制实现(艰难取舍)
//如果没有主键,也需要实现这个方法, return "" 即可
//IEntityStruct 接口的方法,实体类需要实现!!!
func (entity *demoStruct) GetPKColumnName() string {
	//如果没有主键
	//return ""
	return "id"
}

// 1、保存struct对象
func TestInsert(t *testing.T) {

	//需要手动开启事务,匿名函数返回的error如果不是nil,事务就会回滚
	//如果全局DefaultTxOptions配置不满足需求,可以在zorm.Transaction事务方法前设置事务的隔离级别,
	//例如 ctx, _ := dbDao.BindContextTxOptions(ctx, &sql.TxOptions{Isolation: sql.LevelDefault, ReadOnly: false}),如果txOptions为nil,使用全局DefaultTxOptions
	_, err := zorm.Transaction(ctx, func(ctx context.Context) (interface{}, error) {
		//创建一个demo对象
		demo := demoStruct{
			ID:   11,
			Name: "王老九",
			Age:  10,
		}

		//保存对象,参数是对象指针.如果主键是自增,会赋值到对象的主键属性
		_, err := zorm.Insert(ctx, &demo)

		//如果返回的err不是nil,事务就会回滚
		return nil, err
	})
	//标记测试失败
	if err != nil {
		t.Errorf("错误:%v", err)
	}
}

// 2、保存多个Struct对象
func TestInsertSlice(t *testing.T) {

	//需要手动开启事务,匿名函数返回的error如果不是nil,事务就会回滚
	//如果全局DefaultTxOptions配置不满足需求,可以在zorm.Transaction事务方法前设置事务的隔离级别,
	//例如 ctx, _ := dbDao.BindContextTxOptions(ctx, &sql.TxOptions{Isolation: sql.LevelDefault, ReadOnly: false}),如果txOptions为nil,使用全局DefaultTxOptions
	_, err := zorm.Transaction(ctx, func(ctx context.Context) (interface{}, error) {

		//slice存放的类型是zorm.IEntityStruct!!!,golang目前没有泛型,使用IEntityStruct接口,兼容Struct实体类
		demoSlice := make([]zorm.IEntityStruct, 0)

		//创建对象1
		demo1 := demoStruct{
			ID:   12,
			Name: "王老十",
			Age:  9,
		}
		//创建对象2
		demo2 := demoStruct{
			ID:   13,
			Name: "王十一",
			Age:  8,
		}

		demoSlice = append(demoSlice, &demo1, &demo2)

		//批量保存对象,如果主键是自增,无法保存自增的ID到对象里.
		_, err := zorm.InsertSlice(ctx, demoSlice)

		//如果返回的err不是nil,事务就会回滚
		return nil, err
	})
	//标记测试失败
	if err != nil {
		t.Errorf("错误:%v", err)
	}
}

// 3、保存map类型
func TestInsertEntityMap(t *testing.T) {

	//需要手动开启事务,匿名函数返回的error如果不是nil,事务就会回滚
	//如果全局DefaultTxOptions配置不满足需求,可以在zorm.Transaction事务方法前设置事务的隔离级别,
	//例如 ctx, _ := dbDao.BindContextTxOptions(ctx, &sql.TxOptions{Isolation: sql.LevelDefault, ReadOnly: false}),如果txOptions为nil,使用全局DefaultTxOptions
	_, err := zorm.Transaction(ctx, func(ctx context.Context) (interface{}, error) {
		//创建一个EntityMap,需要传入表名
		entityMap := zorm.NewEntityMap("person")
		//设置主键名称
		entityMap.PkColumnName = "id"
		//如果是自增序列,设置序列的值
		//entityMap.PkSequence = "mySequence"

		//Set 设置数据库的字段值
		//如果主键是自增或者序列,不要entityMap.Set主键的值
		entityMap.Set("id", nil)
		entityMap.Set("name", "王十二")
		entityMap.Set("age", "7")

		//执行
		_, err := zorm.InsertEntityMap(ctx, entityMap)

		//如果返回的err不是nil,事务就会回滚
		return nil, err
	})
	//标记测试失败
	if err != nil {
		t.Errorf("错误:%v", err)
	}
}
