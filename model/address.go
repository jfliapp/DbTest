package model

import (
	"fmt"
)

type Address struct{
	Id int `json:"id"`
	Name string `json:"name"`
}

// 查询一个
func GetOneAddress(id string) (Address, error) {
	var mod Address

	// Db.Unsafe() 就是可以结构体跟数据库字段 
	// Get 只能查出一条
	getErr := Db.Unsafe().Get(&mod, "select * from address where id = ?", id)

	if getErr != nil {
		fmt.Println("数据库Get 错误 => ", getErr)
		return mod, getErr
	}

	fmt.Println("数据库Get => ", &mod)

	return mod, getErr
}

// 查询全部
func GetAllAddress() ([]Address, error) {
	var mods []Address

	// Db.Unsafe() 就是可以结构体跟数据库字段 
	// Get 只能查出一条
	getErr := Db.Unsafe().Select(&mods, "select * from address")

	if getErr != nil {
		fmt.Println("数据库Get 错误 => ", getErr)
	}

	fmt.Println("数据库Get => ", &mods)

	return mods, getErr
}

// 删除一个
func DelAddress(id string) error {
	db, err := Db.Unsafe().Exec("delete from address where id = ?", id)
	if err != nil {
		return err
	}
	fmt.Println("删除cheng", db)
	return err
}

// 修改一个
func UpdateAddress(id string, m Address) error {
	db, err := Db.Unsafe().Exec("update address set name=? where id = ?", m.Name, id)
	if err != nil {
		return err
	}
	fmt.Println("修改cheng", db)
	return err
}

// 添加一条数据
func PostAddress(mod Address) error {
	db, err := Db.Unsafe().Exec("insert into address (id, name) value (?, ?)", mod.Id, mod.Name)
	fmt.Println("新增", db)
	return err
}

