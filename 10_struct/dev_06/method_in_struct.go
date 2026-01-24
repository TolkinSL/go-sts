package main

import "fmt"

type CoffeeShop struct {
	nameShop string
}

func (c CoffeeShop) Greet() {
	fmt.Printf("Hello in shop: %s\n", c.nameShop)
}

func main() {
	myShop := CoffeeShop{
		nameShop: "Tula",
	}

	myShop.Greet()
}
