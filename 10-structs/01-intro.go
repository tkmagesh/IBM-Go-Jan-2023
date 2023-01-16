package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

func main() {
	// var emp Employee
	// var emp Employee = Employee{100, "Magesh", 10000}
	// var emp Employee = Employee{Id: 100, Name: "Magesh", Salary: 10000}
	/*
		var emp Employee = Employee{
			Id:     100,
			Name:   "Magesh",
			Salary: 10000,
		}
	*/
	emp := Employee{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
	}
	// fmt.Println(emp)
	// fmt.Printf("%#v\n", emp)
	fmt.Printf("%+v\n", emp)

	fmt.Println("Accessing the fields")
	fmt.Println(emp.Id, emp.Name, emp.Salary)

	e1 := Employee{
		Id:     101,
		Name:   "Kumar",
		Salary: 10000,
	}
	fmt.Printf("&e1 = %p\n", &e1)

	/*
		e2 := Employee{
			Id:     101,
			Name:   "Kumar",
			Salary: 10000,
		}
		fmt.Printf("&e2 = %p\n", &e2)
		fmt.Println(e1 == e2)
	*/

	/*
		e2 := e1
		e2.Salary = 20000
		fmt.Println(e1.Salary, e2.Salary)
	*/

	fmt.Println(FormatEmployee(emp))
	fmt.Println("After awarding 2000 as bonus")
	AwardBonus(&emp, 2000)
	fmt.Println(FormatEmployee(emp))

	// empPtr := &Employee{}
	empPtr := new(Employee) //same as the above
	fmt.Println(empPtr)

	// for Slice, Map, Channel => use make()
}

func FormatEmployee(emp Employee) string {
	return fmt.Sprintf("Id = %d, Name = %q, Salary = %.2f", emp.Id, emp.Name, emp.Salary)
}

func AwardBonus(emp *Employee, bonus float32) {
	// (*emp).Salary += bonus
	emp.Salary += bonus
}
