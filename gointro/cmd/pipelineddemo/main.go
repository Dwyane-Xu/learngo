package main

import (
	"bufio"
	"fmt"
	"learngo/gointro/pipeline"
	"os"
)

func main() {
	file, err := os.Create("large.in")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.RandomSource(100000000)

	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer, p)
	writer.Flush()

	file, err = os.Open("large.in")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p = pipeline.ReaderSource(bufio.NewReader(file), -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}

func mergeDemo() {
	p := pipeline.Merge(pipeline.InMemSort(pipeline.ArraySource(3, 2, 6, 7, 4)),
		pipeline.InMemSort(pipeline.ArraySource(1, 5, 3, 9, 2, 15, 4)))
	for v := range p {
		fmt.Println(v)
	}
}
