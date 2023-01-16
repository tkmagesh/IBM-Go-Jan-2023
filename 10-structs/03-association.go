package main

import "fmt"

type Organization struct {
	Name string
	City string
}

type Employee struct {
	Id     int
	Name   string
	Salary float32
	Org    *Organization
}

func main() {
	ibm := &Organization{
		Name: "IBM",
		City: "Bengaluru",
	}
	e1 := Employee{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
		Org:    ibm,
	}
	fmt.Printf("%+v\n", e1.Org)

	fmt.Println("After changing Org City to Pune")
	ibm.City = "Pune"
	fmt.Printf("%+v\n", e1.Org)
}
