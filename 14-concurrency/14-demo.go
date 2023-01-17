package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go fn(ch)
	for {
		if data, chOpen := <-ch; chOpen {
			fmt.Println(data)
		} else {
			break
		}
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
