package calculator

import "fmt"

var op_count map[string]int

func OpCount() map[string]int {
	return op_count
}

func init() {
	op_count = make(map[string]int)
	fmt.Println("calculator intialized - 1")
}

func init() {
	fmt.Println("calculator intialized - 2")
}
