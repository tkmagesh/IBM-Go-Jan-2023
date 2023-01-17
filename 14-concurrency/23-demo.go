package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 10
	fmt.Println("len(ch) :", len(ch))
	ch <- 20
	fmt.Println("len(ch) :", len(ch))
	data := <-ch
	fmt.Println(data)
	fmt.Println("len(ch) :", len(ch))
	data = <-ch
	fmt.Println(data)
	fmt.Println("len(ch) :", len(ch))

}
