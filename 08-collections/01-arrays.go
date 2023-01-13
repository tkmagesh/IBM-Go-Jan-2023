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

	nosPtr := &nos
	fmt.Println(nos[0], (*nosPtr)[0], nosPtr[0])

	fmt.Println("Before sorting, nos =", nos)
	Sort(&nos)
	fmt.Println("After sorting, nos =", nos)

}

func Sort(nos *[5]int) /* DONOT return anything */ {
	for i := 0; i < len(nos)-1; i++ {
		for j := i + 1; j < len(nos); j++ {
			// if (*nos)[i] > (*nos)[j] {
			if nos[i] > nos[j] {
				nos[i], nos[j] = nos[j], nos[i]
			}
		}
	}
}
