package main

import "fmt"

func main() {
	coffe := "Expresso"
	my_pointer := &coffe

	fmt.Println("Coffe", &coffe)
	fmt.Printf("In Memory coffe %p\n", &coffe)
	fmt.Printf("In Memory coffe %p\n", my_pointer)

	coffeCopy := coffe
	fmt.Printf("In Memory coffeCopy %p\n", &coffeCopy)

	fmt.Printf("In Memory coffe Expresso %p\n", &coffe)
	coffe = "Black"
	fmt.Printf("In Memory coffe Black %p\n", &coffe)

	coffePrice := 4.75
	pointerCoffePrice := &coffePrice
	fmt.Println("CoffePrice", coffePrice)

	*pointerCoffePrice = 6.78
	fmt.Println("CoffePrice", coffePrice)

}