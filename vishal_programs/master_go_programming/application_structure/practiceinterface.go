package main

import "fmt"

type vehicle interface {
	License() string
	Name() string
}

type car2 struct {
	lic   string
	brand string
}

func (c car2) License() string {
	return c.lic
}
func (c car2) Name() string {
	return c.brand
}

func main() {
	var v vehicle = car2{brand: "indica", lic: "WB0D0267"}

	fmt.Println(v.License())
	fmt.Println(v.Name())
}
