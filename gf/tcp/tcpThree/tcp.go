package main

import (
	"fmt"
	"time"

	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"

	"github.com/gogf/gf/net/gtcp"
)

func main() {
	// Server
	go gtcp.NewServer("127.0.0.1:8199", func(conn *gtcp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.Recv(-1)
			if len(data) > 0 {
				if err := conn.Send(append([]byte("> "), data...)); err != nil {
					fmt.Println(err)
				}
			}
			if err != nil {
				break
			}
		}
	}).Run()

	time.Sleep(time.Second)

	// Client
	for {
		if conn, err := gtcp.NewConn("127.0.0.1:8199"); err == nil {
			if b, err := conn.SendRecv([]byte(gtime.Datetime()), -1); err == nil {
				fmt.Println(string(b), conn.LocalAddr(), conn.RemoteAddr())
			} else {
				fmt.Println(err)
			}
			conn.Close()
		} else {
			glog.Error(err)
		}
		time.Sleep(time.Second)
	}
}