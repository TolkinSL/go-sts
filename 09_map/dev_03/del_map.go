package main

import "fmt"

func main() {
	menu := map[string]float64{
		"cola": 2.3,
		"pepsi": 2.1,
		"milk": 3.2,
		"tea": 1.3,
	}

	fmt.Printf("%+v\n", menu)
	
	delete(menu, "milk")
	fmt.Printf("%+v\n", menu)

	delete(menu, "coffee")
	fmt.Printf("%+v\n", menu)

	for i, v := range menu {
		fmt.Printf("drink: %s, price: %0.2f$\n", i, v)
	}
}
