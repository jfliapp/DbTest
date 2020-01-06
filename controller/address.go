package controller

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"Test/dbTest/model"
	"Test/dbTest/server"
)

// 查询单个 根据id
func ControllOneAddress(repw http.ResponseWriter, req *http.Request){
	
	req.ParseForm()

	id := req.Form.Get("id")

	data, err:= server.ServerOneAddress(id)
	if err != nil {
		fmt.Println("查询客户出现问题",err)
		repw.Write([]byte("数据不存在"))
		return
	}
	fmt.Println(data)
	buf, _ := json.Marshal(data)
	// 默认是的数据格式text/plain;我们需要修改它返回的格式为application/json
	repw.Header().Set("Content-Type", "application/json")
	repw.Write(buf)

}
// 查询 全部列表
func ControllAllAddress(repw http.ResponseWriter, req *http.Request){

	data, err:= model.GetAllAddress()
	if err != nil {
		fmt.Println("查询客户出现问题",err)
	}
	fmt.Println(data)
	buf, _ := json.Marshal(data)
	// 默认是的数据格式text/plain;我们需要修改它返回的格式为application/json
	repw.Header().Set("Content-Type", "application/json")
	repw.Write(buf)

}

// 删除单个
func ControllDelAddress(repw http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	id := req.Form.Get("id")

	err:= model.DelAddress(id)
	if err != nil {
		fmt.Println("查询客户出现问题",err)
		repw.Write([]byte("删除shibia"))
		return
	}

	// 默认是的数据格式text/plain;我们需要修改它返回的格式为application/json
	// repw.Header().Set("Content-Type", "application/json")
	repw.Write([]byte("删除成功"))
}

// 修改单个
func ControllUpdateAddress(repw http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	id := req.Form.Get("id")

	buf, _ := ioutil.ReadAll(req.Body)

	var mod model.Address

	json.Unmarshal(buf, &mod)

	err:= model.UpdateAddress(id, mod)
	if err != nil {
		fmt.Println("查询客户出现问题",err)
		repw.Write([]byte("修改shibia"))
		return
	}

	// 默认是的数据格式text/plain;我们需要修改它返回的格式为application/json
	// repw.Header().Set("Content-Type", "application/json")
	repw.Write([]byte("修改成功"))
}

// 添加一个
func ControllPostAddress(repw http.ResponseWriter, req *http.Request) {
	
	var mod model.Address

	d, _ := ioutil.ReadAll(req.Body)

	json.Unmarshal(d, &mod)

	err := server.ServerPostAddress(mod)

	if err != nil {
		fmt.Println("添加cont有啥错",err)
		repw.Write([]byte("fail add"))
		return
	}

	// repw.Header().Set("Content-Type", "application/json")
	// 这里可能要做一个统一的返回格式
	repw.Write([]byte("add success"))
}
