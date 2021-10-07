package main

import "fmt"

type Person struct {
	Name string
	Age  uint
}

func NewPerson(name string, age uint) Person {
	return Person{name, age}
}

type Human struct {
	Age     uint
	Country string
}

func NewHuman(age uint, country string) Human {
	return Human{age, country}
}

func (p Person) Greet() {
	fmt.Println("Hello")
}

type Employee struct {
	Person
	Human
	Salary float64
}

func NewEmployee(name string, age uint, country string, salary float64) Employee {
	return Employee{NewPerson(name, age), NewHuman(age, country), salary}
}

func (e Employee) Payroll() {
	fmt.Println(e.Salary * 0.90)
}

func main() {
	e := NewEmployee("Arthur", 23, "Peru", 5500)
	fmt.Println(e.Name, e.Human.Country)
	e.Payroll()
}
