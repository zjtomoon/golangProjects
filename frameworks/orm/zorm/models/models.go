package models

import (
	"context"
	"fmt"
	"gitee.com/chunanyong/zorm"
)

var db *zorm.DBDao

// 参考： https://blog.csdn.net/ZHOUAXING/article/details/120690265
var ctx = context.Background()

func init() {
	var err error
	dbConfig := zorm.DataSourceConfig{
		//连接数据库DSN
		DSN: "root:password01!@tcp(127.0.0.1:3306)/manage?charset=utf8&parseTime=true",
		//数据库类型
		DriverName: "mysql",
		DBType:     "mysql",
		//是否的打印sql
		PrintSQL: true,
		//最大连接数 默认50
		MaxOpenConns: 0,
		//最大空闲数 默认50
		MaxIdleConns: 0,
		//连接存活秒时间. 默认600
		ConnMaxLifetimeSecond: 0,
		//事务隔离级别的默认配置,默认为nil
		DefaultTxOptions: nil,
	}

	db, err = zorm.NewDBDao(&dbConfig)
	if err != nil {
		fmt.Errorf("数据库连接异常 %v", err)
		panic(err)
	}
	fmt.Println("数据库连接成功")
}
