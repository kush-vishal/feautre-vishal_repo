package main

import "fmt"

type book struct {
	price  float64
	tittle string
}

func (b *book) discount(discount float64) float64 {

	b.price = (100 - discount) * b.price / 100
	return b.price
}

func (b *book) vat(vatPercentage float64) float64 {
	b.price = (100 + vatPercentage) * b.price / 100
	return b.price
}

func (b *book) changePrice(p float64) {
	b.price = p
}

func main() {
	var book1 book
	book1.price = 99.99
	book1.tittle = "oz the great"
	//book1.vat(15)
	//book1.discount(10)
	book1.changePrice(94.876)
	fmt.Printf("\n book price : %v", book1.price)
	fmt.Printf("\n book name : %v", book1.tittle)
	fmt.Printf("\nBook's price:%.2f\n", book1.price)

}
