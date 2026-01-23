package main

import "fmt"

type CoffeeOrder struct {
	CoffeType    string
	CoffeSize    string
	CustomerName string
	BonusPoints  int
}

func main() {
	var order CoffeeOrder
	fmt.Printf("%#v\n", order)

	order.CoffeType = "Cappucino"
	order.CoffeSize = "Large"
	order.CustomerName = "Vasia"
	fmt.Printf("%#v\n", order)
	fmt.Printf("Order size: %s\n", order.CoffeSize)
}
