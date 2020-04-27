package main

import "fmt"

func main() {
	ints := NewIntSet()
	for _, i := range []int{1, 3, 5, 7} {
		ints.Add(i)
		fmt.Println(ints)
	}
	for _, i := range []int{1, 2, 3, 4, 5, 6, 7} {
		fmt.Print(i, ints.Contains(i), " ")
		ints.Delete(i)
		fmt.Println(ints)
	}
	fmt.Println()
	for {
		if err := ints.Undo(); err != nil {
			break
		}
		fmt.Println(ints)
	}
}
