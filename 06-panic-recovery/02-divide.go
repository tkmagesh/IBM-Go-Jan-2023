package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError error = errors.New("divisor cannot be 0")

func main() {
	q, r, e := divideClient(100, 0)
	if e == DivideByZeroError {
		fmt.Println("Do not attempt to divide by 0. try again!")
		return
	}
	if e != nil {
		fmt.Println("unknown error", e)
		fmt.Println("Attempt again")
		return
	}
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)

}

func divideClient(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Divide panicked")
			err = e.(error)
			return
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

//3rd party
func divide(x, y int) (quotient, remainder int) {
	if y == 0 {
		panic(DivideByZeroError)
	}
	fmt.Println("Calculating quotient")
	quotient = x / y
	fmt.Println("Calculating remainder")
	remainder = x % y
	return
}
