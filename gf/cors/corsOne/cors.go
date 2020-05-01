package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
)

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func MidelewareOriginCORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	corsOptions.AllowDomain = []string{"goframe.org", "baidu.com"}
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}

func Order(r *ghttp.Request) {
	r.Response.Write("GET")
}

func main() {
	path := "/Users/xujinzhao/许锦钊/程序/Go/learngo/tmp/glog-cat"
	glog.SetPath(path)
	glog.Stdout(false).Cat("cat1").Cat("cat2").Println("test")
	list, err := gfile.ScanDir(path, "*", true)
	g.Dump(err)
	g.Dump(list)

	/*
		s := g.Server()
		s.Group("/api.v1", func(group *ghttp.RouterGroup) {
			group.Middleware(MidelewareOriginCORS)
			group.GET("/order", Order)
		})
		s.SetPort(8199)
		s.Run()
	*/
}
