package main

import "fmt"

func main() {
	// const pi float32 = 3.14
	fmt.Println("Constants Example")
	/*
		const pi float32 = 3.14
		const app_version string = "1.0"
	*/

	//type inference
	/*
		const pi = 3.14
		const app_version = "1.0"
	*/

	/*
		const (
			pi          float32 = 3.14
			app_version string  = "1.0"
		)
	*/

	/*
		const (
			pi          = 3.14
			app_version = "1.0"
		)
	*/
	const pi, app_version = 3.14, "1.0"
	fmt.Printf("pi = %v, type = %T\n", pi, pi)
	fmt.Printf("app_version = %v, type = %T\n", app_version, app_version)
}
