package main

import "fmt"

func sellDrink(stock map[string]int, drink string, quantity int) {

	_, ok := stock[drink]
	
	fmt.Println("---")
	if ok {
		stock[drink] -= quantity
		fmt.Printf("Sell %s количество %d\n", drink, quantity)
	} else {
		fmt.Printf("%s отсутствует на складе\n", drink)
	}

	fmt.Printf("In func: %+v\n", stock)
	fmt.Printf("In func: %p\n", &stock)
	fmt.Println("---")
}

func main() {
	drinkStock := map[string]int{
		"cola":  14,
		"pepsi": 21,
		"tea":   7,
	}

	fmt.Printf("%+v\n", drinkStock)
	fmt.Printf("%p\n", &drinkStock)

	sellDrink(drinkStock, "pepsi", 5)

	fmt.Printf("%+v\n", drinkStock)
	fmt.Printf("%p\n", &drinkStock)

	sellDrink(drinkStock, "milk", 2)
}