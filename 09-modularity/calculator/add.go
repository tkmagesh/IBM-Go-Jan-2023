package calculator

import "fmt"

func init() {
	fmt.Println("calculator - initialization logic for add")
}

func Add(x, y int) int {
	op_count["add"]++
	return x + y
}
