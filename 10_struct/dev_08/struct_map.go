package main

import "fmt"

type CoffeeOrder struct {
	CustomerName string
	CoffeeType   string
	CoffeeSize   string
	Quantity     int
}

type CoffeMap map[string]string

func (m CoffeMap) printInfo(s string) {
	fmt.Printf("Map my print: %s\n",m[s])
}

func main() {
	//Struct
	orderStr := CoffeeOrder{
		CustomerName: "Vasia",
		CoffeeType:   "Cappuchino",
		CoffeeSize:   "Small",
		Quantity:     2,
	}

	//Map
	var orderMap CoffeMap = map[string]string{
		"CustomerName": "Vasia",
		"CoffeeType":   "Cappuchino",
		"CoffeeSize":   "Small",
	}

	fmt.Printf("Str: %s\n", orderStr.CustomerName)
	fmt.Printf("Str: %s\n", orderStr.CoffeeType)
	fmt.Printf("Map: %s\n", orderMap["CustomerName"])
	fmt.Printf("Map: %s\n", orderMap["CoffeeType"])
}
