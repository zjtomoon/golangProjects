package models

import (
	"context"
	"gitee.com/chunanyong/zorm"
	"testing"
)

// 删除
func TestDelete(t *testing.T) {
	//需要手动开启事务,匿名函数返回的error如果不是nil,事务就会回滚
	//如果全局DefaultTxOptions配置不满足需求,可以在zorm.Transaction事务方法前设置事务的隔离级别,
	//例如 ctx, _ := dbDao.BindContextTxOptions(ctx, &sql.TxOptions{Isolation: sql.LevelDefault, ReadOnly: false}),如果txOptions为nil,使用全局DefaultTxOptions
	_, err := zorm.Transaction(ctx, func(ctx context.Context) (interface{}, error) {
		demo := &demoStruct{}
		demo.ID = 11

		_, err := zorm.Delete(ctx, demo)

		//如果返回的err不是nil,事务就会回滚
		return nil, err
	})
	if err != nil { //标记测试失败
		t.Errorf("错误:%v", err)
	}

}
