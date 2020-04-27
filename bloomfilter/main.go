package bloom

import (
	"fmt"

	"github.com/willf/bloom"
)

func main() {
	n := uint(1000)
	filter := bloom.New(2000*n, 5)
	filter.Add([]byte("Love you"))
	fmt.Println(filter.Test([]byte("Love")))
	//fmt.Println(bloom.EstimateParameters(100000, 0.001))
}
