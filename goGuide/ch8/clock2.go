package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

// 支持传入参数作为端口号
var port = flag.String("port", "8000", "请输入端口")

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleClock2Conn(conn)
	}
}

func handleClock2Conn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
