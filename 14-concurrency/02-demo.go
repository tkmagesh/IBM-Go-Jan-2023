package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1) //increment the counter by 1
	go f1()   //schdeuling this function execution through the scheduler
	f2()
	wg.Wait() // wait for the counter to become 0
}

func f1() {
	defer wg.Done()
	fmt.Println("f1 started")
	time.Sleep(1 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 started")
	fmt.Println("f2 completed")
}
