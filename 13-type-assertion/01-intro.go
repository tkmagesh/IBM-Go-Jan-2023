package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

//fmt.Stringer interface implementation
func (product Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %.2f, Units = %d, Category = %q", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

func main() {
	// var x interface{} //before go1.18
	var x any //introduced in go 1.18

	// x = 100
	// x = "Proident cillum cillum commodo nulla dolor velit consequat pariatur officia nulla ipsum ex proident. Laboris laboris ea excepteur amet veniam pariatur tempor duis magna ad labore culpa adipisicing sit. Veniam ipsum laboris cillum culpa culpa proident qui cillum. Proident in culpa esse amet."
	// x = 99.99
	// x = true
	// x = struct{}{}
	fmt.Println(x)

	// x = 100
	x = "this is a string"
	// y := x.(int) + 200
	// fmt.Println(y)
	if val, ok := x.(int); ok {
		y := val + 200
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	// x = true
	// x = 89.99
	x = Product{}
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x + 200 =", val+200)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case float64:
		fmt.Println("x is a float64, 99% of x = ", val*0.99)
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case struct{}:
		fmt.Println("x is an empty struct")
	case fmt.Stringer:
		fmt.Println("x implements fmt.Stringer(), x.String() =", val.String())
	default:
		fmt.Println("unknown type")
	}

}
