package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DB *sqlx.DB
)

//初始化
func Init(dns string) error {
	var err error
	DB, err = sqlx.Open("mysql", dns)
	if err != nil {
		return err
	}
	//查看是否连接成功
	err = DB.Ping()
	if err != nil {
		return err
	}
	DB.SetMaxOpenConns(100) //设置最大连接
	DB.SetMaxIdleConns(16)  //设置最大空闲
	return nil
}
