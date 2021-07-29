package main
import "fmt"

func temp2() {
	var k int
	var n1, n2 int
	fmt.Print("Please enter First number: ")
	fmt.Scanln(&n1)
	fmt.Print("Please enter Second number: ")
	fmt.Scanln(&n2)
	fmt.Print("Please select  \n1 for add \n2 for subtract \n3 for multiplication \n4 for division\n\n")
	fmt.Scanln(&k)
	result := 0
	switch k {
	case 1:
		result = n1 + n2
		fmt.Println("result : ", result)

	case 2:
		result = n1 - n2
		fmt.Println("result : ", result)
	case 3:
		result = n1 * n2
		fmt.Println("result : ", result)
	case 4:
		result = n1 / n2
		var remainder = n1 % n2
		fmt.Printf("result :%d and remainder :%d", result,remainder)

	default:
		fmt.Println("wrong entry")
	}

}
