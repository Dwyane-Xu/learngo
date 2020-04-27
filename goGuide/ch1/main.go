package main

import "fmt"

func main() {

	var global *int
	var x int
	x = 1
	global = &x
	fmt.Println(*global)

	//FetchAll()
}
