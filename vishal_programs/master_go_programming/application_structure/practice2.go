package main

import "fmt"

var top int
var stack []int

func push(data int) {
	stack[top] = data
	top++

}
func pop() {
	stack[top] = 0

}

func main() {

	var data int
	for {
		var c int
		fmt.Println("enter 1 to push into stack :")
		fmt.Println("\nenter 2 to pop into stack :")
		fmt.Println("\nenter any other to exit :")
		fmt.Scan(&c)
		if c == 1 {
			fmt.Printf("enter the data to push into stack :")
			fmt.Scan(&data)
			push(data)
			fmt.Print(stack)
		} else if c == 2 {

			fmt.Printf("enter the data to push into stack :")
			fmt.Scan(&data)
			pop()
			fmt.Print(stack)
		} else {
			break
		}

	}

}
