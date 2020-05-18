package main

import (
	"bufio"
	"fmt"
	"learngo/gointro/pipeline"
	"os"
	"strconv"
)

func main() {
	p := CreatePipeline("large.in", 800000000, 4)
	WriteToFile(p, "large.out")
	PrintFile("large.out")
}

// WriteToFile 将结果写到文件中
func WriteToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriterSink(writer, p)
}

// PrintFile 打印结果
func PrintFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderSource(file, -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}

// CreatePipeline 首先拆分文件，读取多个文件块到多个channel中，然后进行归并排序
// 返回归并排序的结果channel
func CreatePipeline(filename string, fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount
	pipeline.Init()

	var sortResults []<-chan int
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)

		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)

		sortResults = append(sortResults, pipeline.InMemSort(source))
	}

	return pipeline.MergeN(sortResults...)
}

func CreateNetworkPipeline(filename string, fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount
	pipeline.Init()

	var sortAddr []string
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)

		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)

		addr := ":" + strconv.Itoa(7000+i)
		pipeline.NetworkSink(addr, pipeline.InMemSort(source))
		sortAddr = append(sortAddr, addr)
	}

	var sortResults []<-chan int
	for _, addr := range sortAddr {
		sortResults = append(sortResults, pipeline.NetworkSource(addr))
	}

	return pipeline.MergeN(sortResults...)
}
