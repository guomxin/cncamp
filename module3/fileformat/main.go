package main

import (
	"encoding/json"
	"fmt"
	"log"
	"module3/pkg/apis"
	"os"
)

func main() {
	filePath := "D:\\Git\\cncamp\\module3\\io\\test.txt"
	personalInfo := apis.PersonalInformation{
		Name:   "Guomao",
		Sex:    "男",
		Tall:   1.73,
		Weight: 65,
		Age:    40,
	}

	fmt.Printf("结果是%+v\n", personalInfo)

	data, err := json.Marshal(personalInfo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Marshal的结果是", string(data))

	writeFile(filePath, data)
	readFile(filePath)
}

func readFile(filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("无法读取文件", filePath, ", 错误信息是", err)
		return
	}

	fmt.Println("读取出来的内容是", string(data))

	personalInfo := apis.PersonalInformation{}
	json.Unmarshal(data, &personalInfo)
	fmt.Printf("UnMarshal的结果是%+v\n", personalInfo)
}

func writeFile(filePath string, data []byte) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("无法创建文件", filePath, ", 错误信息是", err)
		return
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("无法写入文件", filePath, ", 错误信息是", err)
		return
	}

}
