package main

import "fmt"

func main() {
	// var productRanks map[string]int
	// var productRanks map[string]int = make(map[string]int)
	// productRanks["pen"] = 4
	// var productRanks map[string]int = map[string]int{"Pen": 4}
	// var productRanks map[string]int = map[string]int{"Pen": 4, "Pencil": 1, "Marker": 2}
	/*
		var productRanks map[string]int = map[string]int{
			"Pen":    4,
			"Pencil": 1,
			"Marker": 2,
		}
	*/
	productRanks := map[string]int{
		"Pen":    4,
		"Pencil": 1,
		"Marker": 2,
	}
	fmt.Println(productRanks)

	productRanks["Scribble-pad"] = 5
	fmt.Println(productRanks)

	fmt.Println("len(productRanks) =", len(productRanks))

	fmt.Println("Iterating a map (using range)")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	fmt.Println("Check for the existence of a key")
	keyToCheck := "Pen"
	// fmt.Println("productRanks[Stappler] =", productRanks[keyToCheck])
	if rank, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q is %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("key [%q] does not exist\n", keyToCheck)
	}

	fmt.Println("Remove an item")
	keyToRemove := "Stappler"
	delete(productRanks, keyToRemove)
	fmt.Println(productRanks)
}
