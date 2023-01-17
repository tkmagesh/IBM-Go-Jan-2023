package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() //schdeuling this function execution through the scheduler
	f2()
	// DO NOT DO THIS
	// time.Sleep(1 * time.Second)
	fmt.Scanln()
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 started")
	fmt.Println("f2 completed")
}
