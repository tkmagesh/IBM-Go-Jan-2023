package main

import "fmt"

func main() {
	defer fmt.Println("	[deferred - main]")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Unknown error:", err)
		}
		fmt.Println("Application exits")
	}()
	fmt.Println("main started")
	f1()
	fmt.Println("main completed")
}

func f1() {
	/*
		defer func() {
			fmt.Println("	[deferred - f1]")
		}()
	*/
	defer fmt.Println("	[deferred - f1] - 1")
	defer fmt.Println("	[deferred - f1] - 2")
	defer fmt.Println("	[deferred - f1] - 3")
	fmt.Println("f1 started")
	f2()
	fmt.Println("f1 completed")
}

func f2() {

	defer func() {
		fmt.Println("	[deferred - f2]")
	}()

	// defer fmt.Println("	[deferred - f2]")
	fmt.Println("f2 started")
	divisor := 7
	_ = 100 / divisor // to panic
	fmt.Println("f2 completed")
}
