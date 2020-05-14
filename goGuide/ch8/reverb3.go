package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func handleReverb3Conn(c net.Conn) {
	input := bufio.NewScanner(c)
	var message = make(chan string)
	var wg sync.WaitGroup

	go func() {
		for {
			select {
			case <-time.After(10 * time.Second):
				c.Close()
			case mes := <-message:
				go func(c net.Conn, shout string, delay time.Duration) {
					defer wg.Done()
					fmt.Fprintln(c, "\t", strings.ToUpper(shout))
					time.Sleep(delay)
					fmt.Fprintln(c, "\t", shout)
					time.Sleep(delay)
					fmt.Fprintln(c, "\t", strings.ToLower(shout))
				}(c, mes, 1*time.Second)
			}
		}
	}()

	for input.Scan() {
		wg.Add(1)
		text := input.Text()
		message <- text
	}
	wg.Wait()
	c.Close()
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Println(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleReverb3Conn(conn)
	}
}
