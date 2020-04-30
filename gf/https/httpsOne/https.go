package main

import "github.com/gogf/gf/net/ghttp"

func main() {
	s := ghttp.GetServer()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writeln("来自于HTTPS的：哈喽世界！")
	})
	s.EnableHTTPS("/Users/xujinzhao/许锦钊/程序/Go/learngo/cert/server.crt",
		"/Users/xujinzhao/许锦钊/程序/Go/learngo/cert/server.key")
	s.SetHTTPSPort(8199)
	s.SetPort(8080)
	s.Run()
}
