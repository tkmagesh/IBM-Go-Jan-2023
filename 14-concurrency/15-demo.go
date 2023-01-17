package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go fn(ch)
	for data := range ch {
		fmt.Println(data)
	}
}

func fn(ch chan int) {
	count := 5
	for i := 1; i <= count; i++ {
		ch <- i * 10
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}
