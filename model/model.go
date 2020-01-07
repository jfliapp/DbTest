package model

import (
	_ "encoding/json"
)

// 请求返回的数据格式
type Resp struct {
	Code int	`json:"code"` // 1 成功 0 失败
	Msg interface{}	`json:"msg"`
	Data interface{}	`json:"data"`
}


// Address 结构体
type Address struct{
	Id int `json:"id"`
	Name string `json:"name"`
}
