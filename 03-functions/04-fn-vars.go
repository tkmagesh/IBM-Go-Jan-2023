/* higher order functions - assign funtions as values to variables */

package main

import "fmt"

type Operation func(int, int) int

func main() {

	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi there!")
	}
	sayHi()

	var greet func(string)
	greet = func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}
	greet("Magesh")

	var getGreetMsg func(string) string
	getGreetMsg = func(userName string) string {
		return fmt.Sprintf("Hi %s, Have a good day!", userName)
	}
	msg := getGreetMsg("Suresh")
	fmt.Println(msg)

	var divide func(int, int) (int, int)
	divide = func(x, y int) (quotient, remainder int) {
		quotient, remainder = x/y, x%y
		return
	}
	q, r := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)

	var op Operation
	op = func(i1, i2 int) int {
		return i1 + i2
	}
	fmt.Println(op(100, 200))

	op = func(i1, i2 int) int {
		return i1 - i2
	}
	fmt.Println(op(100, 200))

}
