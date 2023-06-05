package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

var qa = map[string]string{
	"你好":    "你好",
	"你是谁":   "我是辛豆豆",
	"你是男是女": "你猜猜看",
	"再见":    "再见",
}

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "监听的端口")
	flag.Parse()

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			log.Println("WARNING: 建立连接失败,", err)
			continue
		}
		log.Println(conn)
		go talk(conn)
	}
}

func talk(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		valid, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				time.Sleep(time.Second)
				continue
			}
			log.Println("WARNING: 读取数据失败,", err)
			continue
		}
		content := buf[:valid]
		resp, ok := qa[string(content)]
		if !ok {
			log.Println("没有找到回答")
			conn.Write([]byte("我听不懂"))
			continue
		}

		conn.Write([]byte(resp))
	}
}
