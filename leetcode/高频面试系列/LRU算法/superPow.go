package main

import "fmt"

// 372. 超级次方

const mod = 1337

func superPow(a int, b []int) int {
	if len(b) == 0 {
		return 1
	}

	partOne := myPowTwo(a, b[len(b)-1])
	partTwo := myPowTwo(superPow(a, b[:len(b)-1]), 10)
	return (partOne * partTwo) % mod
}

func myPowOne(a int, k int) int {
	res := 1
	for i := 0; i < k; i++ {
		res *= a
		res %= mod
	}
	return res
}

func myPowTwo(a int, k int) int {
	if k == 0 {
		return 1
	}

	a %= mod
	if k%2 == 1 {
		return (a * myPowTwo(a, k-1)) % mod
	} else {
		sub := myPowTwo(a, k/2)
		return (sub * sub) % mod
	}
}

func main() {
	fmt.Println(superPow(2, []int{3}))
	fmt.Println(superPow(2, []int{1, 0}))
}
