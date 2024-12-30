package main

import "fmt"

func main() {
	type Employee struct {
		ID        int
		FirstName string
		LastName  string
		Address   string
	}
	employee := Employee{LastName: "Doe", FirstName: "John"}
	fmt.Println(employee)
	employee.ID = 1001
	fmt.Println(employee.FirstName)
	// struct pointer altering object - Notice how the struct becomes mutable when you use pointers.
	fmt.Println(employee)
	employeeCopy := &employee
	employeeCopy.FirstName = "David"
	fmt.Println(employee)
}
