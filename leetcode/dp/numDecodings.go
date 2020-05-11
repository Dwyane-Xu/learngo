package main

import "fmt"

// 91. 解码方法

func numDecodingsOne(s string) int {
	if s[0] == '0' {
		return 0
	}

	pre, cur := 1, 1
	for i := 1; i < len(s); i++ {
		tmp := cur
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				cur = pre
			} else {
				return 0
			}
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6') {
			cur = cur + pre
		}
		pre = tmp
	}
	return cur
}

func numDecodingsTwo(s string) int {
	flag := make([]int, len(s))
	for i := range flag {
		flag[i] = -1
	}
	return dfsTwo(s, 0, flag)
}

func dfsTwo(s string, start int, flag []int) int {
	if start == len(s) {
		return 1
	}
	if s[start] == '0' {
		return 0
	}
	if flag[start] != -1 {
		return flag[start]
	}
	ansOne := dfsTwo(s, start+1, flag)
	ansTwo := 0
	if start+1 < len(s) {
		ten := (s[start] - '0') * 10
		one := s[start+1] - '0'
		if ten+one <= 26 {
			ansTwo = dfsTwo(s, start+2, flag)
		}
	}
	flag[start] = ansOne + ansTwo
	return flag[start]
}

func numDecodingsThree(s string) int {
	cur, next := 0, 1
	if s[len(s)-1] != '0' {
		cur = 1
	}
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] == '0' {
			next, cur = cur, 0
			continue
		}
		ten := (s[i] - '0') * 10
		one := s[i+1] - '0'
		if ten+one <= 26 {
			next, cur = cur, cur+next
		} else {
			next = cur
		}
	}
	return cur
}

func main() {
	fmt.Println(numDecodingsThree("226"))
}
