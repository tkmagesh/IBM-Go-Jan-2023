package main

import "fmt"

func main() {
	var i int16 = 100
	var f float32
	f = float32(i) // use the typename like a function for typecasting
	fmt.Println(f)
}
