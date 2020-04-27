package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Dup3() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		CountLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			data, err := ioutil.ReadFile(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				continue
			}

			for _, line := range strings.Split(string(data), "\n") {
				counts[line]++
			}
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
