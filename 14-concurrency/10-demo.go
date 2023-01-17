/* channel behavior */
package main

import (
	"fmt"
	"sync"
)

//results in a deadlock
/*
func main() {
	ch := make(chan int)
	ch <- 100
	data := <-ch
	fmt.Println(data)
}
*/

/*
func main() {
	ch := make(chan int)
	go func() {
		ch <- 100
	}()
	data := <-ch
	fmt.Println(data)
}
*/

//results in a deadlock
/* func main() {
	ch := make(chan int)
	ch <- 100
	go func() {
		data := <-ch
		fmt.Println(data)
	}()
}
*/

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(1)
	go func() {
		data := <-ch
		fmt.Println(data)
		wg.Done()
	}()
	ch <- 100
	wg.Wait()
}
