package controller

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"Test/dbTest/model"
	"Test/dbTest/util"
	"Test/dbTest/server"
)

// 查询单个 根据id
func ControllOneAddress(repw http.ResponseWriter, req *http.Request){
	
	req.ParseForm()

	id := req.Form.Get("id")

	data, err:= server.ServerOneAddress(id)

	repw.Header().Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println("查询客户出现问题",err)
		resp := util.Resp(0, "失败", nil)
		buf, _ := json.Marshal(resp)
		repw.Write(buf)
		return
	}

	resp := util.Resp(1, "成功", data)
	
	buf, _ := json.Marshal(resp)
	// 默认是的数据格式text/plain;我们需要修改它返回的格式为application/json
	repw.Write(buf)

}
// 查询 全部列表
func ControllAllAddress(repw http.ResponseWriter, req *http.Request){

	data, err:= model.GetAllAddress()
	repw.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("查询客户出现问题",err)
		resp := util.Resp(0, "失败", nil)
		buf, _ := json.Marshal(resp)
		repw.Write(buf)
		return
	}
	resp := util.Resp(1, "成功", data)
	buf, _ := json.Marshal(resp)
	// 默认是的数据格式text/plain;我们需要修改它返回的格式为application/json
	repw.Write(buf)

}

// 删除单个
func ControllDelAddress(repw http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	id := req.Form.Get("id")

	err:= model.DelAddress(id)

	repw.Header().Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println("查询客户出现问题",err)
		resp := util.Resp(0, "失败", nil)
		buf, _ := json.Marshal(resp)
		repw.Write(buf)
		return
	}

	// 默认是的数据格式text/plain;我们需要修改它返回的格式为application/json
	// repw.Header().Set("Content-Type", "application/json")
	resp := util.Resp(1, "成功", nil)
	buf, _ := json.Marshal(resp)
	repw.Write(buf)
}

// 修改单个
func ControllUpdateAddress(repw http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	id := req.Form.Get("id")

	buf, _ := ioutil.ReadAll(req.Body)

	var mod model.Address

	json.Unmarshal(buf, &mod)

	err:= model.UpdateAddress(id, mod)

	repw.Header().Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println("查询客户出现问题",err)
		resp := util.Resp(0, "失败", nil)
		buf, _ := json.Marshal(resp)
		repw.Write(buf)
		return
	}

	// 默认是的数据格式text/plain;我们需要修改它返回的格式为application/json
	// repw.Header().Set("Content-Type", "application/json")
	resp := util.Resp(1, "修改成功", nil)
	buf1, _ := json.Marshal(resp)
	repw.Write(buf1)
}

// 添加一个
func ControllPostAddress(repw http.ResponseWriter, req *http.Request) {
	
	var mod model.Address

	d, _ := ioutil.ReadAll(req.Body)

	json.Unmarshal(d, &mod)

	err := server.ServerPostAddress(mod)
	repw.Header().Set("Content-Type", "application/json")


	if err != nil {
		fmt.Println("添加cont有啥错",err)
		resp := util.Resp(0, "失败", nil)
		buf, _ := json.Marshal(resp)
		repw.Write(buf)
		return
	}

	// repw.Header().Set("Content-Type", "application/json")
	// 这里可能要做一个统一的返回格式
	resp := util.Resp(1, "成功", nil)
	buf, _ := json.Marshal(resp)
	repw.Write(buf)
}
