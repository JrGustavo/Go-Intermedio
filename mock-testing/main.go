package main

import (
	"errors"
	"time"
)

type Person struct {
	DNI  string
	Name string
	Age  int
}

type Employee struct {
	ID int
	Person
}

type FullTimeEmployee struct {
	Employee
	endDate string
}

func GetPersonByDNI(dni string) (Person, error) {
	time.Sleep(5 * time.Second)
	// Simulación de búsqueda
	if dni == "" {
		return Person{}, errors.New("DNI not found")
	}
	return Person{DNI: dni, Name: "John Doe", Age: 30}, nil
}

func GetEmployeeById(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	// Simulación de búsqueda
	if id <= 0 {
		return Employee{}, errors.New("Invalid Employee ID")
	}
	return Employee{ID: id, Person: Person{DNI: "1234", Name: "Jane Doe", Age: 25}}, nil
}

func GetFullTimeEmployeeById(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee

	e, err := GetEmployeeById(id)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Employee = e

	p, err := GetPersonByDNI(dni)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Employee.Person = p

	return ftEmployee, nil
}

func main() {
	ftEmployee, err := GetFullTimeEmployeeById(1, "1234")
	if err != nil {
		panic(err)
	}
	println(ftEmployee.Employee.Person.Name)
}
