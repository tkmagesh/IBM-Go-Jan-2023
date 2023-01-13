/*
write a function "genPrimes" that generates the prime numbers from the given start to end and returns the primeNos
the main function receives the prime nos and prints them
*/

package main

import "fmt"

func main() {

	primeNos := genPrimes(3, 100)
	fmt.Println(primeNos)
}

func genPrimes(start, end int) []int {
	result := make([]int, 0)
	for no := start; no <= end; no++ {
		if isPrime(no) {
			result = append(result, no)
		}
	}
	return result
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
