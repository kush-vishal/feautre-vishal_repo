package main

import "fmt"

func main2(){
	a:= []string{"v","s","d","s","k","k","n"}
	b :=[]string{"v","k","l","s","s","n","v"}
	 var c[10]string
	var temp string

	fmt.Println("first slice",a)
	fmt.Println("Second slice",b)
	for t:=range a {
		for i := range b {
			if a[i] == b[t] {
				temp = a[i]
				var f int
				for j := range c {

					if c[j] == temp {
						f++
					}
				}
				if f ==0 {
					c[i] = temp
				}
			}

		}
	}
	fmt.Println(c)
}
