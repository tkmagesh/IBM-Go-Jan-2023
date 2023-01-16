package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

/*
	Create the following utility functions for "Product"
		FormatProduct() => return a formatted string using the fields of the given product
		ApplyDiscount() => update the cost of the given product by applying the given discount (%)
*/

type Dummy struct {
	Name string
}

type PerishableProduct struct {
	Product
	// Name   string  //overriding the Product.Name
	// Dummy
	Expiry string
}

func main() {
	// grapes := PerishableProduct{Product{100, "Grapes", 50, 10, "Food"}, "2 Days"}
	grapes := PerishableProduct{
		Product: Product{
			Id:       100,
			Name:     "Grapes",
			Cost:     50,
			Units:    10,
			Category: "Food",
		},
		Expiry: "2 Days",
	}
	fmt.Printf("%+v\n", grapes)
	fmt.Println(grapes.Expiry)
	fmt.Println(grapes.Product.Name)
	// fmt.Println(grapes.Dummy.Name)
	fmt.Println(grapes.Name)

	/* Use the "FormatProduct()" function to print the grapes (PerishableProduct) */
	/* Use the "ApplyDiscount()" function to apply 10% discount for the grapes and print */
}
