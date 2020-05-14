package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Add(1)
	wg.Done()
	wg.Done()

	wg.Wait()

	wg.Add(1)
	wg.Add(1)

	wg.Wait()

	fmt.Println("111")
}
