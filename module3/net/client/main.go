package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("请求服务端失败:", err)
	}
	defer conn.Close()
	talk(conn, "你好")
}

func talk(conn net.Conn, message string) {
	_, err := conn.Write([]byte(message))
	if err != nil {
		log.Println("发送消息失败:", err)
	} else {
		data := make([]byte, 1024)
		len, err := conn.Read(data)
		if err != nil {
			log.Println("WARNING:读取服务器返回数据时出错", err)
		} else {
			resp := data[:len]
			log.Println("应答:", string(resp))
		}
	}
}
