package main

import "fmt"

func main9() {
	a := []int{2, 3, 43, 24, 23, 45, 65, 12, 3, 44, 55, 6, 9}
	var c int
	fmt.Println(a)

	for i := range a {
		for j:=range  a{
			if a[i]>a[j]{
				c=a[i]
				a[i]=a[j]
				a[j]=c
			}
		}

	}
	fmt.Println(a)
}

