package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// config文件结构
type Configuration struct {
	Address string
	Static  string
}

// 保存json文件解析的数据
var Config Configuration

func LoadConfig(path string) {
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
	}

	// 创建json解析器
	decoder := json.NewDecoder(file)

	// json解析并将数据填入Configuration中
	err = decoder.Decode(&Config)
	if err != nil {
		fmt.Println(err.Error())
	}
}
