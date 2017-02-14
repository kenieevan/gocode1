package main

import "fmt"

type Employee struct {
	ID int
}

var ken Employee

func EmployeeByID() *Employee {
	return &ken
}

func EmployeeByID2() *Employee {
	return &ken
}

func main() {
	EmployeeByID().ID = 123
	fmt.Println(ken.ID)
	ken1 := EmployeeByID2()
	ken1.ID = 12
	fmt.Println(ken1.ID)
	fmt.Println(ken.ID)
}
