package main

import (
	"fmt"
	"strings"
)

/*
Write the apis for the following

IndexOf => return the index of the given product
	ex:  returns the index of the given product

Includes => return true if the given product is present in the collection else return false
	ex:  returns true if the given product is present in the collection

Filter => return a new collection of products that satisfy the given condition
	some of the use cases:
		1. filter all costly products (cost > 1000)
			OR
		2. filter all stationary products (category = "Stationary")
			OR
		3. filter all costly stationary products
		etc

Any => return true if any of the product in the collections satifies the given criteria
	some of the use cases:
		1. are there any costly product (cost > 1000)?
		OR
		2. are there any stationary product (category = "Stationary")?
		OR
		3. are there any costly stationary product?
		etc

All => return true if all the products in the collections satifies the given criteria
	some of use cases:
		1. are all the products costly products (cost > 1000)?
		OR
		2. are all the products stationary products (category = "Stationary")?
		OR
		3. are all the products costly stationary products?
		etc
*/

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (product Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %.2f, Units = %d, Category = %q", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

type Products []Product

func (products Products) Format() string {
	var builder strings.Builder
	for _, p := range products {
		builder.WriteString(fmt.Sprintf("%s\n", p.Format()))
	}
	return builder.String()
}

func (products Products) IndexOf(p Product) int {
	for idx, product := range products {
		if p == product {
			return idx
		}
	}
	return -1
}

func (products Products) FilterByCost(cost float32) Products {
	result := Products{}
	for _, p := range products {
		if p.Cost >= cost {
			result = append(result, p)
		}
	}
	return result
}

func (products Products) FilterByCategory(category string) Products {
	result := Products{}
	for _, p := range products {
		if p.Category == category {
			result = append(result, p)
		}
	}
	return result
}

func (products Products) Filter(predicate func(Product) bool) Products {
	result := Products{}
	for _, p := range products {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}

func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}

	kettle := Product{101, "Kettle", 2500, 10, "Utencil"}
	fmt.Println("Index of kettle : ", products.IndexOf(kettle))

	fmt.Println("Initial List")
	fmt.Println(products.Format())

	fmt.Println("Costly Products [cost > 1000]")
	// costlyProducts := products.FilterByCost(1000)
	costlyProductPredicate := func(p Product) bool {
		return p.Cost >= 1000
	}
	costlyProducts := products.Filter(costlyProductPredicate)
	fmt.Println(costlyProducts.Format())

	fmt.Println("Stationary Products [category = stationary]")
	// stationaryProducts := products.FilterByCategory("Stationary")
	stationaryProductPredicate := func(p Product) bool {
		return p.Category == "Stationary"
	}
	stationaryProducts := products.Filter(stationaryProductPredicate)
	fmt.Println(stationaryProducts.Format())
}
