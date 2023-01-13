package main

import "fmt"

func main() {
	// var nos []int
	// var nos []int = []int{3, 1, 4, 2, 5}
	// var nos []int = []int{3, 1, 4}
	// var nos = []int{3, 1, 4}
	nos := []int{3, 1, 4, 2, 5}
	fmt.Printf("&nos = %p, nos = %v\n", &nos, nos)

	/* appending new values */
	nos = append(nos, 10)
	fmt.Printf("&nos = %p, nos = %v\n", &nos, nos)

	nos = append(nos, 20, 30, 40)
	fmt.Printf("&nos = %p, nos = %v\n", &nos, nos)

	hundreds := []int{100, 200, 300}
	nos = append(nos, hundreds...) // hundres[0], hundreds[1]

	fmt.Println("Iterating an Slice (using indexer)")
	for i := 0; i < len(nos); i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iterating an Slice (using range)")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	x := []int{10, 20, 30, 40, 50}
	dupX := x
	dupX[0] = 100
	fmt.Println("x =", x)
	fmt.Println("dupX =", dupX)

	fmt.Println("After sorting")
	Sort(nos)
	fmt.Println(nos)

	fmt.Println("slicing")
	fmt.Println("nos[2:6] =", nos[2:6]) //=> values from index 2 to (6-1)
	fmt.Println("nos[2:] =", nos[2:])   // values from index 2 to the end
	fmt.Println("nos[:6] =", nos[:6])   // values from start to (6-1)

	subset := nos[2:6]
	subset[0] = 1000
	fmt.Println(nos)
	fmt.Printf("len(subset) = %d, subset = %v\n", len(subset), subset)
}

func Sort(nos []int) /* DONOT return anything */ {
	for i := 0; i < len(nos)-1; i++ {
		for j := i + 1; j < len(nos); j++ {
			// if (*nos)[i] > (*nos)[j] {
			if nos[i] > nos[j] {
				nos[i], nos[j] = nos[j], nos[i]
			}
		}
	}
}
