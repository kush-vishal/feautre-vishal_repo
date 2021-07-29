package main

import "fmt"

type par struct {
	name string
	id   int
}

var slice2 = make([]int, 0, 10)
var slice3 = make([]string, 0, 10)
var slice4 = make([]par, 0, 10)

func main() {
	slice1 := []int{1, 2, 3, 4}
	fmt.Println("slice practice")
	fmt.Println("slice1 : ", slice1)
	slice3 = []string{"random1", "random3", "random4"}
	fmt.Println(slice3)
	slice4 = []par{{"random", 3}, {"random2", 4}, {"random10", 10}}
	fmt.Println(slice4)
	fmt.Println(slice4[1:])
	fmt.Println(slice4[0:])

}
