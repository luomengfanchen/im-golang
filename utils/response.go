package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 封装的josn结构体
type H map[string]interface{}

// 响应json数据
func ResponseJSON(w http.ResponseWriter, code int, obj interface{}) {
	// 设置响应头，返回json数据
	w.Header().Set("Content-Type","application/json")
	// 返回状态码
	w.WriteHeader(code)

	// 解析结构体返回字符串
	ret, err := json.Marshal(obj)
	if err != nil {
		fmt.Println(err.Error())
	}

	// 将响应数据返回
	w.Write(ret)
}