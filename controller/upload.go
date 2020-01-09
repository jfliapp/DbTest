package controller

import (
	"fmt"
	"net/http"
	// "io"
	"io/ioutil"
	"os"
)

const (
	UPLOAD_DIR = "./static"
)


func UploadImg(repW http.ResponseWriter, req *http.Request) {

	req.ParseForm()

	req.ParseMultipartForm(32 << 20) // 这个是用来获取 图片文件的

	fileHeader := req.MultipartForm.File["file"][0]
	b := req.Form.Get("text")
	fmt.Println("req => ", b)

	filename := fileHeader.Filename
	fmt.Println("文件名上传的=>", filename)

	file, err := fileHeader.Open() // 他自己的方法
	defer file.Close()
	if err == nil {
			// data, err := ioutil.ReadAll(file)
				t, _ := os.Create(UPLOAD_DIR+"/"+filename)
				
				data, _ := ioutil.ReadAll(file)

				n, _ := t.Write(data)

				fmt.Println(n, "n1, n 读了多少行")

				defer t.Close()
	}
		
	repW.Header().Set("Content-Type","application/json")
	repW.Write([]byte("{\"s\":\"sss\"}"))
}
