package main

import (
	"fmt"
	"log"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	time.Sleep(5 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}

func main() {
	fmt.Println(triple(4))
}
