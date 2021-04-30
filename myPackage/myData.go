package myPackage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var (
	Num0 int = 7
)

type decodeData struct {
	//必须大写名字,否则无法读取
	MyId   int    `json:"id"`
	MyName string `json:"name"`
}

func PrintJsonData() {
	jsonFile := "data/myJson.json"
	file, err := os.Open(jsonFile)
	if err != nil {
		log.Fatal("open file error")
	}

	//延迟关闭文件
	defer file.Close()

	var tempData []*decodeData
	//json文件读取
	json.NewDecoder(file).Decode(&tempData)

	for _, v := range tempData {
		fmt.Println(v)
	}
}
