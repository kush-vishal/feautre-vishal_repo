package main

import "fmt"

func main4() {
	a := []string{"v", "s", "d", "s", "k", "k", "n"}
	d := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "w", "x", "y", "z"}
	var c = make([]int, 26)

	fmt.Println(" slice : ", a)
	for i:= range d{
		for j:=range a{
			if a[j]==d[i]{
				c[i]=c[i]+1
			}
		}
	}
	for i:=range d{
		if c[i]>0{
			fmt.Printf("\nfrequency of %v is %v",d[i],c[i])
		}
	}
}