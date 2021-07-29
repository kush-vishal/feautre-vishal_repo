package main

import (
	"fmt"
	"strings"
)

func main(){
	var sentence string
	//var word string
	s:="my name is vishal"
	fmt.Println(" Hi Please enter sentence to find word")
	//fmt.Scan(&sentence)
	fmt.Println(sentence)
	fmt.Println(s)
	fmt.Println("Enter that you want to find : ")
	//fmt.Scan(&word)
	splitWord := strings.Split(s, " ")
	fmt.Println(splitWord)


}
