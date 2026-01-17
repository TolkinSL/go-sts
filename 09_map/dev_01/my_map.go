package main

import "fmt"

func main() {
	menu := map[string]float64{
		"cola": 2.5,
		"pepsi": 2.3,
		"tea": 1.6,
		"milk": 3.2,
	}

	fmt.Printf("%+v\n", menu)
	fmt.Println(menu["tea"])

	menu["pepsi"] = 2.1
	fmt.Printf("%+v\n", menu)

	val, ok := menu["pepsi"]
	fmt.Println(val, ok)

	val, ok = menu["coffee"]
	fmt.Println("coffee", val, ok)

	fmt.Println("Menu length", len(menu))

	menu["coffee"] = 2.7
	val, ok = menu["coffee"]
	fmt.Println("coffee", val, ok)

	fmt.Println("Menu length", len(menu))
}