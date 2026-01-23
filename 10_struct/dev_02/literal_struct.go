package main

import "fmt"

type CoffeeOrder struct {
	CoffeeType string
	CoffeeSize string
}

func main() {
	order := CoffeeOrder{
		CoffeeType: "Latte",
		CoffeeSize: "Small",
	}
	fmt.Printf("Type: %s\n", order.CoffeeType)
	fmt.Println()
	
	order2 := CoffeeOrder{
		"Cappucino",
		"Big",
	}
	fmt.Printf("%#v\n", order2)
}
