package main

import "fmt"

func main() {
	fn := getFn(1)
	exec(fn)

	fn = getFn(2)
	exec(fn)
}

func getFn(id int) func() {
	if id == 1 {
		return f1
	}
	return f2
}

func exec(fn func()) {
	fn()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
