package main

import (
	"fmt"

	"github.com/fatih/color"
	calc "github.com/tkmagesh/ibm-go-jan-2023/09-modularity/calculator" // importing with an alias
	"github.com/tkmagesh/ibm-go-jan-2023/09-modularity/calculator/utils"
	// _ "github.com/tkmagesh/ibm-go-jan-2023/09-modularity/calculator" // importing ONLY to execute the "init" function(s)
)

func main() {
	// fmt.Println("module executed")
	color.Red("module executed")
	/*
		fmt.Println(calculator.Add(100, 200))
		fmt.Println(calculator.Subtract(100, 200))
		// fmt.Println(calculator.op_count)
		fmt.Println(calculator.OpCount())
	*/

	fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
	// fmt.Println(calc.op_count)
	fmt.Println(calc.OpCount())

	fmt.Println("Is 21 even? :", utils.IsEven(21))

}
