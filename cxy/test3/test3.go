package main

func isHappy(n int) bool {
	st := make(map[int]bool, 0)
	for !st[n] {
		st[n] = true
		next := 0
		for n != 0 {
			next += (n % 10) * (n % 10)
			n /= 10
		}
		n = next
		if n == 1 {
			return true
		}
	}
	return false
}

func main() {
	//fmt.Println(isIsomorphic("aba", "baa"))
}
