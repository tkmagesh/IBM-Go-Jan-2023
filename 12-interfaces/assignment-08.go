package main

import (
	"fmt"
	"sort"
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

func (product Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %.2f, Units = %d, Category = %q", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

type Products []Product

func (products Products) String() string {
	var builder strings.Builder
	for _, p := range products {
		builder.WriteString(fmt.Sprintf("%s\n", p))
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

func (products Products) Filter(predicate func(Product) bool) Products {
	result := Products{}
	for _, p := range products {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}

/* sort.Interface implementation */
func (products Products) Len() int {
	return len(products)
}

// Default comparison implementation (by Id)
func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

// sort by Name
type ByName struct {
	Products
}

func (byName ByName) Less(i, j int) bool {
	return byName.Products[i].Name < byName.Products[j].Name
}

// sort by Cost
type ByCost struct {
	Products
}

func (byCost ByCost) Less(i, j int) bool {
	return byCost.Products[i].Cost < byCost.Products[j].Cost
}

// sort by Units
type ByUnits struct {
	Products
}

func (byUnits ByUnits) Less(i, j int) bool {
	return byUnits.Products[i].Units < byUnits.Products[j].Units
}

// sort by Category
type ByCategory struct {
	Products
}

func (byCategory ByCategory) Less(i, j int) bool {
	return byCategory.Products[i].Category < byCategory.Products[j].Category
}

// Utility method for sorting
func (products Products) Sort(attrName string) {
	switch attrName {
	case "Id":
		sort.Sort(products)
	case "Name":
		sort.Sort(ByName{products})
	case "Cost":
		sort.Sort(ByCost{products})
	case "Units":
		sort.Sort(ByUnits{products})
	case "Category":
		sort.Sort(ByCategory{products})
	default:
		sort.Sort(products)
	}
}

// alternative sort implementation using sort.Slice()
func (products Products) SortNew(attrName string, desc bool) {
	getDescComparer := func(fn func(i, j int) bool) func(i, j int) bool {
		return func(i, j int) bool {
			return !fn(i, j)
		}
	}
	switch attrName {
	case "Id":
		comparer := func(i, j int) bool {
			return products[i].Id < products[j].Id
		}
		if desc {
			comparer = getDescComparer(comparer)
		}
		sort.Slice(products, comparer)
	case "Name":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Name < products[j].Name
		})
	case "Cost":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Cost < products[j].Cost
		})
	case "Units":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Units < products[j].Units
		})
	case "Category":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Category < products[j].Category
		})
	}
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

	fmt.Println("Initial List")
	fmt.Println(products)

	// Default sort (by Id)
	fmt.Println("Sort by Id (default)")
	// sort.Sort(products)
	// products.Sort("Id")
	products.SortNew("Id", false)
	fmt.Println(products)

	fmt.Println("Sort by Id (descending) (default)")
	// sort.Sort(products)
	// products.Sort("Id")
	products.SortNew("Id", true)
	fmt.Println(products)

	//sort by Name
	fmt.Println("Sort by Name")
	// sort.Sort(ByName{products})
	// products.Sort("Name")
	products.SortNew("Name", false)
	fmt.Println(products)

	//sort by Cost
	fmt.Println("Sort by Cost")
	// sort.Sort(ByCost{products})
	// products.Sort("Cost")
	products.SortNew("Cost", false)
	fmt.Println(products)

	//sort by Units
	fmt.Println("Sort by Units")
	// sort.Sort(ByUnits{products})
	// products.Sort("Units")
	products.SortNew("Units", false)
	fmt.Println(products)

	//sort by Category
	fmt.Println("Sort by Category")
	// sort.Sort(ByCategory{products})
	// products.Sort("Category")
	products.SortNew("Category", false)
	fmt.Println(products)
}
