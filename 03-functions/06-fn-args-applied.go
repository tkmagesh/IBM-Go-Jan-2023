/* higher order functions - functions as arguments  */
package main

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
)

func main() {
	/*
		log.Println("Operation Started...")
		add(100, 200)
		log.Println("Operation Completed...")

		log.Println("Operation Started...")
		subtract(100, 200)
		log.Println("Operation Completed...")
	*/

	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/
	logOperation(add, 100, 200)
	logOperation(subtract, 100, 200)
}

func logOperation(operation func(int, int), x, y int) {
	funcName := runtime.FuncForPC(reflect.ValueOf(operation).Pointer()).Name()
	log.Printf("Operation[%q] started with %d and %d...", funcName, x, y)
	operation(x, y)
	log.Println("Operation completed...")
}

/* ============== */

func logAdd(x, y int) {
	log.Printf("Operation started with %d and %d...", x, y)
	add(x, y)
	log.Println("Operation completed...")
}

func logSubtract(x, y int) {
	log.Printf("Operation started with %d and %d...", x, y)
	subtract(x, y)
	log.Println("Operation completed...")
}

//3rd party library - we cannot change the code
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}

/*
func add(x, y int) {
	log.Println("Operation Started...")
	fmt.Println("Add Result :", x+y)
	log.Println("Operation Completed...")
}

func subtract(x, y int) {
	log.Println("Operation Started...")
	fmt.Println("Subtract Result :", x-y)
	log.Println("Operation Completed...")
}
*/
