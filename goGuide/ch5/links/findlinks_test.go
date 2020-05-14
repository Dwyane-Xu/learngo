package links_test

import (
	"fmt"
	"testing"

	"learngo/goGuide/ch5/links"
)

func TestFinklinks(t *testing.T) {
	input := make([]string, 1)
	input[0] = "http://www.baidu.com"
	fmt.Println(links.Crwal(input[0]))
}
