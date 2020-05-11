package main

import (
	"fmt"
	"runtime"
	"time"
)

func gen3(nums ...int) <-chan int {
	out := make(chan int)
	// make buffer size to len(nums)
	// out := make(chan int, len(nums))
	for _, n := range nums {
		go func(n int) {
			fmt.Println("begin", n)
			out <- n
			fmt.Println("done", n)
		}(n)
	}
	return out
}

func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()

	out := gen3(2, 3, 4)

	for n := range out {
		fmt.Println(n)
		time.Sleep(2 * time.Second)
		if true {
			break
		}
	}
}
