package main

import "fmt"

func main() {
	var no int
	no = 100

	var noPtr *int
	// address of the value
	noPtr = &no

	fmt.Println(no, noPtr)

	//value at an address (dereferencing)
	val := *noPtr
	fmt.Println(val)

	fmt.Println(no == *(&no))

	fmt.Println("Before incrementing, no =", no)
	increment(&no)
	fmt.Println("After incrementing, no =", no)

	n1, n2 := 100, 200
	fmt.Printf("Before swapping, n1 = %d and n2 = %d\n", n1, n2)
	swap( /*  */ )
	fmt.Printf("Before swapping, n1 = %d and n2 = %d\n", n1, n2)
}

func increment(x *int) {
	fmt.Println("x =", x)
	// x++
	*x++
}

func swap( /*  */ ) {

}
