package main

import (
	"fmt"
	"sync"
)

func makeThumbnails6(strs ...string) int {
	length := make(chan int)
	var wg sync.WaitGroup
	for _, s := range strs {
		wg.Add(1)
		// worker
		go func(s string) {
			defer wg.Done()
			fmt.Println(s)

			length <- len([]rune(s))
		}(s)
	}

	// closer
	go func() {
		wg.Wait()
		close(length)
	}()

	var total int
	for leng := range length {
		total += leng
	}
	return total
}

func main() {
	fmt.Println(makeThumbnails6("asdf", "w4t", "12312312"))
}
