/*
write a function "genPrimes" that generates the prime numbers from the given start to end and returns the primeNos
the main function receives the prime nos and prints them
*/

package main

import "fmt"

func main() {

LOOP:
	for no := 3; no <= 100; no++ {
		for i := 2; i <= (no / 2); i++ {
			if no%i == 0 {
				continue LOOP
			}
		}
		fmt.Println("Prime No : ", no)
	}
}
