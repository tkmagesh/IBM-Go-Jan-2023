package main

import (
	"fmt"
	"sync"
)

// share memory by communicating (preferred)

func main() {
	wg := &sync.WaitGroup{}
	result := 0
	wg.Add(1)
	go add(100, 200, wg, &result)
	wg.Wait()
	fmt.Println("Result :", result)
}

func add(x, y int, wg *sync.WaitGroup, result *int) {
	*result = x + y
	wg.Done()
}
