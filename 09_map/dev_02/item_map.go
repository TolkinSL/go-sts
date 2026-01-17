package main

import "fmt"

func main() {
	menu := map[string]float64{
		"cola":  2.4,
		"pepsi": 2.1,
	}

	drink := "tea"

	fmt.Println(drink, menu[drink])

	val, ok := menu[drink]
	fmt.Println(drink, val, ok)

	var my_menu1 map[string]float64
	fmt.Printf("%#v\n", my_menu1)

	val, ok = my_menu1[drink]
	fmt.Println(drink, val, ok)

}
