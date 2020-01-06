package main

import (
	 // 这个是它默认的用来操作数据库
	"database/sql"
	"fmt"

	// 安装方式: go get github.com/go-sql-driver/mysql
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbhostsip  = "127.0.0.1:3306"
	dbusername = "root"
	dbpassword = ""
	dbname     = "goTest"
)

// 初始化数据库
func InitDB() (*sql.DB, error) {
	//打开数据库，前面是驱动名称，所以要导入:mysql驱动github.com/go-sql-driver/mysql
	db, err := sql.Open("mysql", dbusername+":"+dbpassword+"@"+"tcp("+dbhostsip+")/"+dbname)
	if nil != err {
		fmt.Println("Open Database Error:", err)
		return nil, err
	}
	// 设置数据库的最大连接数
	db.SetConnMaxLifetime(100)
	// 设置数据库最大的闲置连接数
	db.SetMaxIdleConns(10)
	// 验证连接
	if err = db.Ping(); nil != err {
		fmt.Println("Open Database Fail,Error:", err)
		return nil, err
	}
	fmt.Println("Connect Success!!!")
	defer db.Close()

	return db, nil
}

func main(){
	Db, err := InitDB()
	if err != nil{
		fmt.Println("cuowu")
	}
	fmt.Println("数据库实例", Db)
}