package main

import (
	"fmt"
)

type emp struct {
	empid int
}

func main() {
	data := findEmp(1)
	if data != nil {
		fmt.Println(*data)
	} else {
		fmt.Println("got nil")
	}

}

func findEmp(empid int) *emp {
	var employee emp
	if empid == 12 {
		employee.empid = 12
		return &employee
	}
	return nil

}
