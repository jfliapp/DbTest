package main

import (
	"fmt"
	"net/http"	
	"Test/dbTest/controller"
)


func main() {

	// 这里有问题 要在里面判断 请求的方法是POST GET DELETE
	http.HandleFunc("/list", controller.ControllAllAddress)
	http.HandleFunc("/", controller.ControllOneAddress) // 查询单个
	http.HandleFunc("/del", controller.ControllDelAddress) // 删除单个
	http.HandleFunc("/update", controller.ControllUpdateAddress) // 修改单个
	http.HandleFunc("/addAddress", controller.ControllPostAddress) // 添加
	http.ListenAndServe(":8888", nil)

	fmt.Print("主程序 main.go")
}
