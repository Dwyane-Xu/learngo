package main

import (
	"fmt"
	"time"
)

func main() {
	//var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				fmt.Printf("This is %d\n", i)
				//a[i]++
			}
		}(i)
	}
	time.Sleep(time.Minute)
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("%d is %d\n", i, a[i])
	//}
}
