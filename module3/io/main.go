package main

import (
	"fmt"
	"os"
)

func main() {
	// filePath := "D:\\Git\\cncamp\\module3\\io\\test.txt"
	// file, err := os.Open(filePath)
	// if err != nil {
	// 	fmt.Println("无法打开文件", filePath, ", 错误信息是", err)
	// 	os.Exit(1)
	// }
	// defer file.Close()

	// b := make([]byte, 1024)
	// n, err := file.Read(b)
	// if err != nil {
	// 	fmt.Println("无法读取文件", err)
	// 	os.Exit(1)
	// }

	// fmt.Println("读出的内容", string(b[:n]))
	// fmt.Println("读出的大小", n)

	filePath := "D:\\Git\\cncamp\\module3\\io\\test2.txt"
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("无法创建文件", filePath, ", 错误信息是", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString("豆豆喜欢和爸爸一起玩")
}
