package main

import "fmt"

var av = map[int]string{ 10:"Mr. stark", 20:" Mr. tony stark"}

type team struct {
	av3 map[int]string

}

func temp() {
	var av2 = map[int]string{}
	av2[1]="Mr. Steve"
	av2[2]="Mr. Steve Rogers"
	var avenger team
	fmt.Println(av2)
	delete(av2,1)
	fmt.Println(av)
	fmt.Println(av2)
	fmt.Println(avenger.av3)
}
