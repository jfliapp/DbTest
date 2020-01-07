package model

import (
	"fmt"
	// 自己下插件来操作数据库 他默认的操作数据库是 "database/sql"---》操作数据比较麻烦
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sqlx.DB

func init() {
	db, err := sqlx.Open(`mysql`, `root:@tcp(localhost:3306)/goTest?charset=utf8&parseTime=true`)

	if err != nil {
		fmt.Print(err.Error())
	}
	dberr := db.Ping()
	if dberr != nil {
		fmt.Print("数据库连接错误", dberr)
	}

	Db = db
}
