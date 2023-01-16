package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

/* Convert the following into methods of Product */
func FormatProduct(product Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %.2f, Units = %d, Category = %q", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

func ApplyDiscount(product *Product, discount float32) {
	product.Cost = product.Cost * ((100 - discount) / 100)
}

type PerishableProduct struct {
	Product
	Expiry string
}

func main() {

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

	/* Use the "FormatProduct()" function to print the grapes (PerishableProduct) */
	fmt.Println(FormatProduct(grapes.Product))
	/* Use the "ApplyDiscount()" function to apply 10% discount for the grapes and print */
	ApplyDiscount(&grapes.Product, 10)
	fmt.Println(FormatProduct(grapes.Product))
}
