/* higher order functions - functions as arguments  */
package main

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"time"
)

type OperationFn func(int, int)
type CrossCuttingConcernFn func(OperationFn) OperationFn

func main() {

	logAdd := logOperation(add)
	profileLogAdd := profileOpertation(logAdd)
	profileLogAdd(100, 200)

	logSubtract := logOperation(subtract)
	profileLogSubtract := profileOpertation(logSubtract)
	profileLogSubtract(100, 200)
}

func chain(operation OperationFn, crossCuttingConcerns ...CrossCuttingConcernFn) OperationFn {
	for i := 0; i < len(crossCuttingConcerns); i++ {
		crossCuttingConcern := crossCuttingConcerns[i]
		operation = crossCuttingConcern(operation)
	}
	return operation
}

func profileOpertation(operation OperationFn) OperationFn {
	return func(x, y int) {
		funcName := runtime.FuncForPC(reflect.ValueOf(operation).Pointer()).Name()
		start := time.Now()
		operation(x, y)
		elapsed := time.Since(start)
		log.Printf("%q took %v\n", funcName, elapsed)
	}
}

func logOperation(operation OperationFn) OperationFn {
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
