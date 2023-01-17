package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go fn(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func fn(ch chan int) {
	ch <- 10
	time.Sleep(1 * time.Second)
	ch <- 20
	time.Sleep(1 * time.Second)
	ch <- 30
	time.Sleep(1 * time.Second)
	ch <- 40
	time.Sleep(1 * time.Second)
	ch <- 50
	time.Sleep(1 * time.Second)
}
