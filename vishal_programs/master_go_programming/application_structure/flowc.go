package main

import "fmt"

func div(number, divider int) {
	for i := 1; i <= number; i++ {
		if i%divider == 0 {
			fmt.Printf("\n %d number is divisible by %d", i, divider)

		}

	}
}
func div2(number, divider, limit int) {
	for i := 1; i <= number; i++ {
		var c int
		if i%divider != 0 {

			continue
		} else {
			fmt.Printf("\n %d number is divisible by %d", i, divider)
			c++
			if c == limit {
				break
			}

		}

	}
}

func div3(number, divider, limit int) {
	var c int
	for i := 1; i <= number; i++ {

		if i%divider == 0 {
			fmt.Printf("\n %d number is divisible by %d", i, divider)
			c++
			if c == limit {
				break
			}

		}
	}
}
func div4(number, divider1, divider2 int) {
	for i := 1; i <= number; i++ {
		if i%divider1 == 0 && i%divider2 == 0 {
			fmt.Printf("\n %d number is divisible by %d and %d", i, divider1, divider2)

		}

	}
}

func main() {
	var (
		numbe    int
		divider  int
		limit    int
		divider2 int
	)
	fmt.Println("please enter number : ")
	fmt.Scan(&numbe)
	fmt.Printf("enter divider number from 1 to 100: ")
	fmt.Scan(&divider)
	fmt.Printf("enter 2nd divider number from 1 to 100: ")
	fmt.Scan(&divider2)
	fmt.Printf("enter limit : ")
	fmt.Scan(&limit)
	//div(numbe, divider)
	//div3(numbe, divider, limit)
	div4(numbe, divider, divider2)
}
