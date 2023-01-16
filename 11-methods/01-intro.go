package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

// as a function
/*
func FormatEmployee(emp Employee) string {
	return fmt.Sprintf("Id = %d, Name = %q, Salary = %.2f", emp.Id, emp.Name, emp.Salary)
}
*/

//as a method
func (this Employee) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Salary = %.2f", this.Id, this.Name, this.Salary)
}

//as a function
/*
func AwardBonus(emp *Employee, bonus float32) {
	// (*emp).Salary += bonus
	emp.Salary += bonus
}
*/

//as a method
func (this *Employee) AwardBonus(bonus float32) {
	this.Salary += bonus
}

func main() {

	emp := Employee{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
	}

	// fmt.Println(FormatEmployee(emp))
	fmt.Println(emp.Format())
	fmt.Println("After awarding 2000 as bonus")
	// AwardBonus(&emp, 2000)
	emp.AwardBonus(2000)
	// fmt.Println(FormatEmployee(emp))
	fmt.Println(emp.Format())

}
