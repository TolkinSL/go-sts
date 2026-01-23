package main

import "fmt"

// Передача в функции по значению (копия):
// Базовые типы: int, float, bool
// string
// Array
// Struct
// Pointer

type CoffeeOrder struct {
	CoffeeType string
	CoffeeSize string
}

func printOrder(order CoffeeOrder) {
	fmt.Printf("Func: %#v\n", order)
	order.CoffeeSize = "Small"
	fmt.Printf("Func: %#v\n", order)
	fmt.Println()
}

func main() {
	order := CoffeeOrder{
		CoffeeType: "Cappuchino",
		CoffeeSize: "Large",
	}

	printOrder(order)
	fmt.Printf("Main: %#v\n", order)
}
