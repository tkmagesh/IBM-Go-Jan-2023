package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError error = errors.New("divisor cannot be 0")

func main() {
	q, r, err := divide(100, 0)
	if err == DivideByZeroError {
		fmt.Println("Please donot attempt to divide by 0")
		return
	}
	if err != nil {
		fmt.Println("unknown error - ", err)
		return
	}
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)
}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		err := fmt.Errorf("invalid arguments. divisor = %d", y)
		return 0, 0, err
	}
	fmt.Println("Calculating quotient")
	quotient := x / y
	fmt.Println("Calculating remainder")
	remainder := x % y
	return quotient, remainder, nil
}
*/

/*
func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		// err = fmt.Errorf("invalid arguments. divisor = %d", y)
		err = errors.New("divisor cannot be 0")
		return
	}
	fmt.Println("Calculating quotient")
	quotient = x / y
	fmt.Println("Calculating remainder")
	remainder = x % y
	return
}
*/

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		// err = fmt.Errorf("invalid arguments. divisor = %d", y)
		err = DivideByZeroError
		return
	}
	fmt.Println("Calculating quotient")
	quotient = x / y
	fmt.Println("Calculating remainder")
	remainder = x % y
	return
}
