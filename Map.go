package main

import "fmt"

type Employee struct {
	Salary  int
	Country string
}

func main() {
	// Creating and initializing a map
	employeeSalary := make(map[string]int)
	fmt.Println("Empty map:", employeeSalary)

	// Adding items to a map
	employeeSalary["steve"] = 12000
	employeeSalary["jamie"] = 15000
	employeeSalary["mike"] = 9000
	fmt.Println("Employee salaries:", employeeSalary)

	// Initializing a map during declaration
	employeeSalary = map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	employeeSalary["mike"] = 9000
	fmt.Println("Updated employee salaries:", employeeSalary)

	// Retrieving value for a key
	employee := "jamie"
	salary := employeeSalary[employee]
	fmt.Println("Salary of", employee, "is", salary)

	// Checking if a key exists
	newEmp := "joe"
	value, ok := employeeSalary[newEmp]
	if ok {
		fmt.Println("Salary of", newEmp, "is", value)
	} else {
		fmt.Println(newEmp, "not found")
	}

	// Iterating over all elements in a map
	fmt.Println("Contents of the map")
	for key, value := range employeeSalary {
		fmt.Printf("employeeSalary[%s] = %d\n", key, value)
	}

	// Deleting items from a map
	delete(employeeSalary, "steve")
	fmt.Println("Map after deletion:", employeeSalary)

	// Map of structs
	employeeInfo := map[string]Employee{
		"Steve": {12000, "USA"},
		"Jamie": {14000, "Canada"},
		"Mike":  {13000, "India"},
	}
	fmt.Println("Employee information:")
	for name, info := range employeeInfo {
		fmt.Printf("Employee: %s Salary:$%d  Country: %s\n", name, info.Salary, info.Country)
	}

	// Length of the map
	fmt.Println("Length of the map:", len(employeeSalary))

	// Maps are reference types
	modified := employeeSalary
	modified["mike"] = 18000
	fmt.Println("Employee salary changed", employeeSalary)

	// Maps equality
	// Maps can't be compared using the == operator
	// Only nil comparison is allowed
	map1 := map[string]int{"one": 1, "two": 2}
	map2 := map1
	if map1 == nil && map2 == nil {
		fmt.Println("Both maps are nil.")
	}
}
