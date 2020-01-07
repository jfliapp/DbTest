package util

import (
	"Test/dbTest/model"
)

func Resp(code int, msg string, data interface{}) (res model.Resp) {
	res.Code = code
	res.Msg = msg
	res.Data = data

	return res
}
