package main

import (
	"io"
	"log"
	"net"
	"os"
)

func mustCopy3(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Println(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		// 如果TCP的读连接关闭会报错：use of closed network connection
		_, err := io.Copy(os.Stdout, conn)
		log.Println(err)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy3(conn, os.Stdin)
	cw := conn.(*net.TCPConn)
	cw.CloseWrite()
	<-done
}
