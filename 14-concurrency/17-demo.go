package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		// go PrintNo(i, wg)
		go func(no int) {
			fmt.Println(no)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func PrintNo(no int, wg *sync.WaitGroup) {
	fmt.Println(no)
	wg.Done()
}
