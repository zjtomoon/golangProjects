package models

import (
	"context"
	"gitee.com/chunanyong/zorm"
	"testing"
)

// 1、更新struct对象,更新不为零值的字段.主键必须有值
func TestUpdateNotZeroValue(t *testing.T) {

	//需要手动开启事务,匿名函数返回的error如果不是nil,事务就会回滚
	//如果全局DefaultTxOptions配置不满足需求,可以在zorm.Transaction事务方法前设置事务的隔离级别,
	//例如 ctx, _ := dbDao.BindContextTxOptions(ctx, &sql.TxOptions{Isolation: sql.LevelDefault, ReadOnly: false}),如果txOptions为nil,使用全局DefaultTxOptions
	_, err := zorm.Transaction(ctx, func(ctx context.Context) (interface{}, error) {
		//声明一个对象的指针,用于更新数据
		demo := &demoStruct{}
		demo.ID = 1
		demo.Name = "王大"

		_, err := zorm.UpdateNotZeroValue(ctx, demo)

		//如果返回的err不是nil,事务就会回滚
		return nil, err
	})
	if err != nil { //标记测试失败
		t.Errorf("错误:%v", err)
	}

}

// 2、更新struct对象,更新所有字段.主键必须有值
func TestUpdate(t *testing.T) {

	//需要手动开启事务,匿名函数返回的error如果不是nil,事务就会回滚
	//如果全局DefaultTxOptions配置不满足需求,可以在zorm.Transaction事务方法前设置事务的隔离级别,
	//例如 ctx, _ := dbDao.BindContextTxOptions(ctx, &sql.TxOptions{Isolation: sql.LevelDefault, ReadOnly: false}),如果txOptions为nil,使用全局DefaultTxOptions
	_, err := zorm.Transaction(ctx, func(ctx context.Context) (interface{}, error) {

		//声明一个对象的指针,用于更新数据
		demo := &demoStruct{}
		demo.ID = 1
		demo.Name = "王老大"

		_, err := zorm.Update(ctx, demo)

		//如果返回的err不是nil,事务就会回滚
		return nil, err
	})
	//age会为0
	if err != nil { //标记测试失败
		t.Errorf("错误:%v", err)
	}
}

// 3、通过finder更新,zorm最灵活的方式,可以编写任何更新语句,甚至手动编写insert语句
func TestUpdateFinder(t *testing.T) {
	//需要手动开启事务,匿名函数返回的error如果不是nil,事务就会回滚
	//如果全局DefaultTxOptions配置不满足需求,可以在zorm.Transaction事务方法前设置事务的隔离级别,
	//例如 ctx, _ := dbDao.BindContextTxOptions(ctx, &sql.TxOptions{Isolation: sql.LevelDefault, ReadOnly: false}),如果txOptions为nil,使用全局DefaultTxOptions
	_, err := zorm.Transaction(ctx, func(ctx context.Context) (interface{}, error) {
		finder := zorm.NewUpdateFinder("person") // UPDATE person SET
		//finder = zorm.NewDeleteFinder(demoStructTableName)  // DELETE FROM t_demo
		//finder = zorm.NewFinder().Append("UPDATE").Append(demoStructTableName).Append("SET") // UPDATE t_demo SET
		finder.Append("name=?,age=?", "王大", 100).Append("WHERE id=?", 1)

		//更新 "sql":"UPDATE t_demo SET  name=?,age=? WHERE id=?","args":["王大",100,"1"]
		_, err := zorm.UpdateFinder(ctx, finder)

		//如果返回的err不是nil,事务就会回滚
		return nil, err
	})
	if err != nil { //标记测试失败
		t.Errorf("错误:%v", err)
	}

}

// 4、更新一个EntityMap,主键必须有值
func TestUpdateEntityMap(t *testing.T) {
	//需要手动开启事务,匿名函数返回的error如果不是nil,事务就会回滚
	//如果全局DefaultTxOptions配置不满足需求,可以在zorm.Transaction事务方法前设置事务的隔离级别,
	//例如 ctx, _ := dbDao.BindContextTxOptions(ctx, &sql.TxOptions{Isolation: sql.LevelDefault, ReadOnly: false}),如果txOptions为nil,使用全局DefaultTxOptions
	_, err := zorm.Transaction(ctx, func(ctx context.Context) (interface{}, error) {
		//创建一个EntityMap,需要传入表名
		entityMap := zorm.NewEntityMap("person")
		//设置主键名称
		entityMap.PkColumnName = "id"
		//Set 设置数据库的字段值,主键必须有值
		entityMap.Set("id", "1")
		entityMap.Set("name", "王老大")
		//更新 "sql":"UPDATE person SET name=? WHERE id=?","args":["王老大","1"]
		_, err := zorm.UpdateEntityMap(ctx, entityMap)

		//如果返回的err不是nil,事务就会回滚
		return nil, err
	})
	if err != nil { //标记测试失败
		t.Errorf("错误:%v", err)
	}

}
