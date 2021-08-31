package utils

import (
	"database/sql"
	"testing"
	_ "github.com/go-sql-driver/mysql"
)

func Test_GetOne(t *testing.T)  {
	db,err := sql.Open("mysql","root:123.abc@tco(localhost:3306)/test")
	defer func() {
		db.Close()
	}()
	if err != nil {
		t.Fatal(err)
	}

	//测试empty
	car_brand,err := GetOne(db,"select * from user where id = 999999")
	if (car_brand != nil) || (err != nil) {
		t.Fatal("empty测试错误")
	}
}
