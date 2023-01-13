package main

import "fmt"

func main() {
	// var nos [5]int
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}
	// var nos [5]int = [5]int{3, 1, 4}
	// var nos = [5]int{3, 1, 4}
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("Iterating an Array (using indexer)")
	for i := 0; i < 5; i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iterating an Array (using range)")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	x := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("%p\n", &x)
	y := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("%p\n", &y)
	fmt.Println(x == y)

	dupX := x
	dupX[0] = 100
	fmt.Println("x =", x)
	fmt.Println("dupX =", dupX)

	fmt.Println("Before sorting, nos =", nos)
	Sort(?)
	fmt.Println("After sorting, nos =", nos)

}

func Sort(/*  */)/* DONOT return anything */ {

}