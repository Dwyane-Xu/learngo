package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	index := 0
	for _, v := range t {
		if uint8(v) == s[index] {
			index++
		}
		if index == len(s) {
			return true
		}
	}
	return false
}

func isSubsequenceForBig(s string, t string) bool {
	tMap := make(map[byte][]int, 26)
	for i := 0; i < len(t); i++ {
		bt := t[i] - 'a'
		if tMap[bt] == nil {
			tMap[bt] = make([]int, 0)
		}
		tMap[bt] = append(tMap[bt], i)
	}

	index := -1
	for i := 0; i < len(s); i++ {
		idMap := tMap[s[i]-'a']
		if idMap == nil || idMap[len(idMap)-1] <= index {
			return false
		}

		if idMap[0] > index {
			index = idMap[0]
			continue
		}

		l, r := 0, len(idMap)
		for l < r {
			m := (l + r) / 2
			if idMap[m] > index {
				r = m
			} else {
				l = m + 1
			}
		}
		if l == len(idMap) {
			return false
		}
		index = idMap[l]
	}

	return true
}

func main() {
	fmt.Println(isSubsequenceForBig("abc", "ahbgdc"))
	fmt.Println(isSubsequenceForBig("axc", "ahbgdc"))
}
