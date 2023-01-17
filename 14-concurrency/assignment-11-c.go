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
		data := []struct {
			start int
			end   int
		}{
			{start: 3, end: 100},
			{start: 101, end: 200},
			{start: 201, end: 300},
		}
		for _, d := range data {
			wg.Add(1)
			go func(start, end int) {
				ch := genPrimes(start, end)
				for primeNo := range ch {
					resultCh <- primeNo
				}
				wg.Done()
			}(d.start, d.end)
		}

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
