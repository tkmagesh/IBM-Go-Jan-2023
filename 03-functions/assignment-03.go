/*
refactor the following implementations using functions
IMPORTANT : Make sure each function follows SRP (Single Responsibility Principle)
*/

package main

import "fmt"

var userChoice, n1, n2, result int

func main() {

	for {
		userChoice = getUserChoice()
		if userChoice == 5 {
			break
		}

		if userChoice < 1 || userChoice > 5 {
			fmt.Println("Invalid choice")
			continue
		}
		op := getOperation(userChoice)
		n1, n2 = getOperands()
		result = op(n1, n2)
		fmt.Println("Result :", result)

	}
}

func getUserChoice() int {
	var userChoice int
	fmt.Println("Enter your choice:")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Scanln(&userChoice)
	return userChoice
}

func getOperands() (n1, n2 int) {
	fmt.Println("Enter the values:")
	fmt.Scanln(&n1, &n2)
	return
}

func getOperation(userChoice int) func(int, int) int {
	switch userChoice {
	case 1:
		return add
	case 2:
		return subtract
	case 3:
		return multiply
	case 4:
		return divide
	default:
		return func(i1, i2 int) int {
			return 0
		}
	}
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}
