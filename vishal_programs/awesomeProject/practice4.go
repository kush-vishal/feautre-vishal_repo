package main

import "fmt"

func main6(){
	a:=[]int{2,3,43,24,23,45,65,12,3,44,55,6,9}
	b:=[]int{32,45,23,62,1,76,4,67,0}
	fmt.Println(a)
	fmt.Println(b)
	var max int
	for i:=range a{
		if a[i]>max{
			max=a[i]
		}

	}
	for i:=range b{
		if b[i]>max{
			max=b[i]
		}
	}
	fmt.Println(max)
}
