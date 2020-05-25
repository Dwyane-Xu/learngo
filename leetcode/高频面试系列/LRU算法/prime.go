package main

// 素数筛选

func prime(n int) []bool {
	isPrim := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrim[i] = true
	}

	for i := 2; i*i < n; i++ {
		if isPrim[i] {
			for j := i * i; j < n; j += i {
				isPrim[j] = false
			}
		}
	}

	return isPrim
}

func main() {

}
