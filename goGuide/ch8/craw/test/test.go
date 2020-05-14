package main

import (
	"flag"
	"fmt"
	"log"

	"learngo/goGuide/ch5/links"
)

// 统计一共爬取了多个页面
var count = make(chan int)

// 令牌 tokens 是一个计数信号量
// 确保并发请求限制在 20 个以内
var tokens = make(chan struct{}, 20)

func crawl(url string, depth int) urllist {
	fmt.Println(depth, <-count, url)
	// 获取令牌
	tokens <- struct{}{}
	list, err := links.Extract(url)
	// 释放令牌
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return urllist{list, depth + 1}
}

var depth int

func init() {
	// 小于0就是不限制递归深度，0就是只爬取当前页面
	flag.IntVar(&depth, "depth", -1, "深度限制")
}

type urllist struct {
	urls  []string
	depth int
}

func main() {
	// 负责 count 值自增的 goroutine
	go func() {
		var i int
		for {
			i++
			count <- i
		}
	}()

	flag.Parse()
	worklist := make(chan urllist)
	// 等待发送到任务列表的数量
	// 因为需要在 goroutine 里修改，需要换成并发安全的计数器
	// var n sync.WaitGroup
	starturls := flag.Args()
	if len(flag.Args()) == 0 {
		starturls = []string{"http://lab.scrapyd.cn/"}
	}

	var n int
	n++
	go func() { worklist <- urllist{starturls, 0} }()

	// 并发爬取 Web
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list.urls {
			if !seen[link] && (depth >= 0 && list.depth <= depth) {
				seen[link] = true
				n++
				go func(url string, listDepth int) {
					worklist <- crawl(url, listDepth)
				}(link, list.depth)
			}
		}
	}
}
