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
	resultCh := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}

		wg.Add(1)
		go func() {
			ch := genPrimes(3, 100)
			for primeNo := range ch {
				resultCh <- primeNo
			}
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			ch := genPrimes(101, 200)
			for primeNo := range ch {
				resultCh <- primeNo
			}
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			ch := genPrimes(201, 300)
			for primeNo := range ch {
				resultCh <- primeNo
			}
			wg.Done()
		}()

		wg.Wait()
		close(resultCh)
	}()

	for primeNo := range resultCh {
		fmt.Println(primeNo)
	}

}

func genPrimes(start, end int) <-chan int {
	ch := make(chan int)
	go func() {
		for no := start; no <= end; no++ {
			if isPrime(no) {
				ch <- no
				time.Sleep(500 * time.Millisecond)
			}
		}
		close(ch)
	}()
	return ch
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
