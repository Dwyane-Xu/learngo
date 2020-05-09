package main

import "fmt"

// 面试题 08.01. 三步问题

func waysToStep(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else if n == 3 {
		return 4
	}
	one, two, three, mo := 1, 2, 4, int(1e9+7)
	for i := 3; i < n; i++ {
		one, two, three = two, three, (one+two+three)%mo
	}
	return three
}

func main() {
	fmt.Println(waysToStep(4))
}
