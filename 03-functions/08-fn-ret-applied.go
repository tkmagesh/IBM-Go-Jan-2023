/* higher order functions - functions as arguments  */
package main

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"time"
)

func main() {

	/*
		logOperation(add, 100, 200)
		logOperation(subtract, 100, 200)
	*/

	/*
		logAdd := logOperation(add)
		logAdd(100, 200)

		logSubtract := logOperation(subtract)
		logSubtract(100, 200)
	*/

	/*
		profileAdd := profileOpertation(add)
		profileAdd(100, 200)

		profileSubtract := profileOpertation(subtract)
		profileSubtract(100, 200)
	*/

	/*
		logAdd := logOperation(add)
		profileLogAdd := profileOpertation(logAdd)
		profileLogAdd(100, 200)

		logSubtract := logOperation(subtract)
		profileLogSubtract := profileOpertation(logSubtract)
		profileLogSubtract(100, 200)
	*/

	/*
		profileLogAdd := profileOpertation(logOperation(add))
		profileLogAdd(100, 200)

		profileLogSubtract := profileOpertation(logOperation(subtract))
		profileLogSubtract(100, 200)
	*/

	profileLogAdd := chain(add, logOperation, profileOpertation)
	profileLogAdd(100, 200)

	profileLogSubtract := chain(subtract, logOperation, profileOpertation)
	profileLogSubtract(100, 200)

}

func chain(operation func(int, int), crossCuttingConcerns ...func(func(int, int)) func(int, int)) func(int, int) {
	for i := 0; i < len(crossCuttingConcerns); i++ {
		crossCuttingConcern := crossCuttingConcerns[i]
		operation = crossCuttingConcern(operation)
	}
	return operation
}

func profileOpertation(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		funcName := runtime.FuncForPC(reflect.ValueOf(operation).Pointer()).Name()
		start := time.Now()
		operation(x, y)
		elapsed := time.Since(start)
		log.Printf("%q took %v\n", funcName, elapsed)
	}
}

func logOperation(operation func(int, int)) func(int, int) {
	return func(x, y int) {
		funcName := runtime.FuncForPC(reflect.ValueOf(operation).Pointer()).Name()
		log.Printf("Operation[%q] started with %d and %d...", funcName, x, y)
		operation(x, y)
		log.Println("Operation completed...")
	}
}

//3rd party library - we cannot change the code
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
