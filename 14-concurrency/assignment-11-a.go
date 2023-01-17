/*
Modify the below program 5 goroutines are employed to generate prime numbers between 3 to 500
g1 = 3 - 100
g2 = 101 - 200
g3 = 201 - 300
etc
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(1)
	go genPrimes(3, 100, ch, wg)

	wg.Add(1)
	go genPrimes(101, 200, ch, wg)

	wg.Add(1)
	go genPrimes(201, 300, ch, wg)

	go func() {
		for primeNo := range ch {
			fmt.Println(primeNo)
		}
	}()

	wg.Wait()
}

func genPrimes(start, end int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for no := start; no <= end; no++ {
		if isPrime(no) {
			ch <- no
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
