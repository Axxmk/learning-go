package main

import "fmt"

type Employee interface {
	GetName() string
	GetSalary() uint
}

// * Engineer
type Engineer struct {
	Name   string
	Salary uint
}

// Implement methods match the function signature
// defined within Employee interface

func (e *Engineer) GetName() string {
	return "Engineer Name: " + e.Name
}

func (e *Engineer) GetSalary() uint {
	return e.Salary
}

// * Manager
type Manager struct {
	Name   string
	Salary uint
}

func (m *Manager) GetName() string {
	return "Manager Name: " + m.Name
}

func (m *Manager) GetSalary() uint {
	return m.Salary
}

// PrintDetails is a value receiver method that implements the Employee interface
func PrintDetails(e Employee) {
	fmt.Println(e.GetName())
	fmt.Printf("Salary: %d\n", e.GetSalary())
}

func main() {
	engineer := &Engineer{
		Name:   "Ann",
		Salary: 300000,
	}

	manager := &Manager{
		Name:   "Tom",
		Salary: 500000,
	}

	PrintDetails(engineer)
	PrintDetails(manager)
}
