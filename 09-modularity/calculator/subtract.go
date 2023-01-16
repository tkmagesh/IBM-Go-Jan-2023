package calculator

import "fmt"

func init() {
	fmt.Println("calculator - initialization logic for subtract")
}

func Subtract(x, y int) int {
	op_count["subtract"]++
	return x - y
}
