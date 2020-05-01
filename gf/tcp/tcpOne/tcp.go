package main

import (
	"fmt"
	"time"

	"github.com/gogf/gf/os/glog"

	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/net/gtcp"
)

func main() {
	// Server
	go gtcp.NewServer("127.0.0.1:8199", func(conn *gtcp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.Recv(-1)
			if len(data) > 0 {
				fmt.Println(string(data))
			}
			if err != nil {
				break
			}
		}
	}).Run()

	time.Sleep(time.Second)

	// Client
	conn, err := gtcp.NewConn("127.0.0.1:8199")
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10000; i++ {
		if err := conn.Send([]byte(gconv.String(i))); err != nil {
			glog.Error(err)
		}
		time.Sleep(time.Second)
	}
}
