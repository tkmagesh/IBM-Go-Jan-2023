package main

import (
	"fmt"
	"strings"
)

/*
Sort => to sort the products by any attribute
	IMPORTANT: use sort.Sort() function
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
