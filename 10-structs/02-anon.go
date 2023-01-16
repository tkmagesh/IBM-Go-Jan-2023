package main

import "fmt"

func main() {
	/*
		emp := Employee{
			Id:     100,
			Name:   "Magesh",
			Salary: 10000,
		}
	*/

	emp := struct {
		Id     int
		Name   string
		Salary float32
	}{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
	}
	// fmt.Printf("id = %d, Name = %q\n", emp.Id, emp.Name)
	fmt.Println(FormatEmployee(emp))

	dataObj := struct{}{}
	fmt.Println(dataObj)
}

func FormatEmployee(emp struct {
	Id     int
	Name   string
	Salary float32
}) string {
	return fmt.Sprintf("id = %d, Name = %q", emp.Id, emp.Name)
}
