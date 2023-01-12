package main

import "fmt"

func main() {
	sayHi()
	greet("Magesh")
	fmt.Println(getGreetMsg("Suresh"))
	fmt.Println(add(100, 200))

	/* using multiple return results */
	// fmt.Println(divide(100, 7))

	/*
		q, r := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)
	*/

	//Use ONLY the quotient
	q, _ := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d\n", q)

	_, r := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, remainder = %d\n", r)

	_, _ = divide(100, 7) // equivalent to divide(100,7)
}

/* No parameters and No return results */
func sayHi() {
	fmt.Println("Hi there!")
}

/* 1 Parameter */
func greet(userName string) {
	fmt.Printf("Hi %s, Have a nice day!\n", userName)
}

/* 1 Parameter and 1 return result */
func getGreetMsg(userName string) string {
	return fmt.Sprintf("Hi %s, Have a good day!", userName)
}

/* 2 Parameters and 1 return result */
/*
func add(x int, y int) int {
	return x + y
}
*/
func add(x, y int) int {
	return x + y
}

/* 2 Parameters and 2 return results */
/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

//using named results
/*
func divide(x, y int) (quotient int, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
*/

func divide(x, y int) (quotient, remainder int) {
	/* quotient & remainder are initialized */
	quotient, remainder = x/y, x%y
	return
}
