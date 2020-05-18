package main

import "fmt"

func f() (r int) {
	defer func(r int) {
		r = r + 5
		fmt.Println(r)
	}(r)
	return 1
}

func main() {
	fmt.Println(f())
}
