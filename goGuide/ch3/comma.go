package main

import (
	"bytes"
	"strings"
)

func CommaOne(s string) string {
	n := len(s)
	if n < 3 {
		return s
	}
	return CommaOne(s[:n-3]) + "," + s[n-3:]
}

func CommaTwo(s string) string {
	var buf bytes.Buffer
	rest := len(s) % 3
	buf.WriteString(s[0:rest] + ",")
	for i := rest; i < len(s)-3; i += 3 {
		buf.WriteString(s[i:i+3] + ",")
	}
	buf.WriteString(s[len(s)-3:])
	return buf.String()
}

func Same(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for _, v := range s1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}

	return true
}
