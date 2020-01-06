package server

import (
	"fmt"
	"Test/dbTest/model"
)

// 查询一个
func ServerOneAddress(id string) (model.Address, error) {

	data, err:= model.GetOneAddress(id)

	if err != nil {
		fmt.Println("查询客户出现问题",err)		
	}
	
	return data, err
}

// 添加一个
func ServerPostAddress(mod model.Address) (error) {

	data, err := ServerOneAddress(string(mod.Id))

	if data.Id == mod.Id {
		fmt.Println("这条数据已经存在 请添加新的id")
		return err
	}

	err1 := model.PostAddress(mod)

	if err1 != nil {
		fmt.Println("添加cont有啥错",err)
	}
	return err1
}