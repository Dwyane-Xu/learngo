package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"learngo/goGuide/ch5/links"
)

func crwal(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	unseenLinke := make(chan string)
	var wg sync.WaitGroup
	var wg2 sync.WaitGroup

	wg.Add(1)
	go func() { worklist <- os.Args[1:] }()
	go func() {
		wg.Wait()
		close(worklist)
	}()

	for i := 0; i < 10; i++ {
		go func() {
			for link := range unseenLinke {
				foundLinks := crwal(link)
				wg.Add(1)
				wg2.Done()
				// avoid deadline
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				wg2.Add(1)
				unseenLinke <- link
			}
		}
		wg2.Wait()
		wg.Done()
	}
}
