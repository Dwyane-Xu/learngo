package main

import "fmt"

func main() {
	var s []int

	for i := 1; i < 100; i++ {
		fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	s2 := make([]int, 16)
	copy(s2, s1)
	fmt.Println(s2)
	s2 = append(s2[:3], s2[4:]...)
	fmt.Println(s2)

	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	fmt.Println(s2)

	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	fmt.Println(s2)
}
