package main

import "fmt"

// Estructuras básicas
type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

// FullTimeEmployee con método getMessage
type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full-time employee"
}

// Interface PrintInfo
type PrintInfo interface {
	getMessage() string
}

// Función que usa PrintInfo
func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

// TemporaryEmployee con método getMessage
type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

func (tEmployee TemporaryEmployee) getMessage() string {
	return "Temporary employee"
}

// main
func main() {
	// Crear un FullTimeEmployee
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Name"
	ftEmployee.age = 22
	ftEmployee.id = 1
	fmt.Printf("%v\n", ftEmployee)

	// Crear un TemporaryEmployee
	tEmployee := TemporaryEmployee{}
	tEmployee.name = "TempName"
	tEmployee.age = 25
	tEmployee.id = 2

	// Usar la función getMessage
	getMessage(tEmployee)
	getMessage(ftEmployee)
}
