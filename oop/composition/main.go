package main

import (
	"fmt"
	"math/rand"
)

// basic details of the employee
type Employee struct {
	id      int
	name    string
	dept    string
	address Address
}

// address of the employee
type Address struct {
	city string
	zip  int
}

var companyDirectory map[int]Employee

func main() {

	companyDirectory = make(map[int]Employee)

	// address of employee anbu
	address := Address{
		city: "chennai",
		zip:  600119,
	}

	// employee details
	anbu := Employee{
		id:      201911121,
		name:    "anbu",
		dept:    "finance",
		address: address,
	}

	// employee details
	// for the sake of demonstration lets keep same address of another employee
	siva := Employee{
		id:      201911121,
		name:    "siva",
		dept:    "security",
		address: address,
	}

	companyDirectory[rand.Intn(1000)] = anbu
	companyDirectory[rand.Intn(1000)] = siva

	printEmployeeRecords(companyDirectory)
}

func printEmployeeRecords(records map[int]Employee) {
	// iterate over company directory and print employee details
	for id, info := range records {
		fmt.Println("Id: ", id,
			" Name: ", info.name,
			" Department: ", info.dept,
			" Address: ", info.address.city, " ",
			info.address.zip)
	}
}
