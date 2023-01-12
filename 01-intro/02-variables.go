package main

import "fmt"

// var msg string = "Hello World [pkg]!"

var msg = "Hello World [pkg]!"

// msg := "Hello World [pkg]!" //=> NOT ALLOWED

func main() {
	/*
		fmt.Println("Hello World!")
	*/

	/*
		var msg string
		msg = "Hello World!"
	*/

	/*
		var msg string = "Hello World!"
	*/

	//type inference
	/*
		var msg = "Hello World!"
	*/

	// msg := "Hello World"
	/*
		NOTE:
			":=" syntax is valid ONLY in a function and NOT in package scope
	*/
	fmt.Println(msg)

	// Multiple variables
	/*
		var x int
		var y int
		var str string
		var result int
	*/

	/*
		var x, y, result int
		var str string
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	//declaration & initialization

	/*
		var x int = 100
		var y int = 200
		var str string = "Sum of %d and %d is %d\n"
		var result int = x + y
	*/

	/*
		var x, y int = 100, 200
		var str string = "Sum of %d and %d is %d\n"
		var result int = x + y
	*/

	/*
		var (
			x, y   int    = 100, 200
			str    string = "Sum of %d and %d is %d\n"
			result int    = x + y
		)
	*/

	//type inference
	/*
		var (
			x, y   = 100, 200
			str    = "Sum of %d and %d is %d\n"
			result = x + y
		)
	*/

	/*
		var (
			x, y, str = 100, 200, "Sum of %d and %d is %d\n"
			result    = x + y
		)
	*/

	// using ":="

	x, y, str := 100, 200, "Sum of %d and %d is %d\n"
	result := x + y

	fmt.Printf(str, x, y, result)
}
