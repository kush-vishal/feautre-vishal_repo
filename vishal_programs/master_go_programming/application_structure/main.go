package main

import (
	"fmt"
)

var i int

type person struct {
	name string
	age  int
	sex  string
}

func createPerson(age int, name string, sex string) person {

	var p person
	p.name = name
	p.age = age

	return p
}

func names(nameFamily ...person) {

	familyGroup := make(map[int]person)
	for i = range nameFamily {
		familyGroup[i] = nameFamily[i]

	}
	fmt.Print(familyGroup)

}

func main() {

	var endp = make([]person, 0, 10)
	var newMember person
	var k int
	for {
		fmt.Println("Enter 1 to add new member  :")
		fmt.Println("Enter any other number  to exit :")
		fmt.Scan(&k)
		if k == 1 {
			fmt.Println("Enter the name :")
			fmt.Scan(&newMember.name)
			fmt.Println("Enter age of the person :")
			fmt.Scan(&newMember.age)
			fmt.Println("TYPE M For male and F for female :")
			fmt.Scan(&newMember.sex)
			endp = append(endp, newMember)
		} else {
			break
		}
	}
	names(endp[:]...)

	for i = range endp {
		fmt.Printf("\n\nnew members : %v", endp[i])

	}

}
