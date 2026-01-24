package main

import "fmt"

type DrinkType string

func (d DrinkType) selectDrink() {
	fmt.Printf("%#v\n", d)
	fmt.Printf("I select: %s\n", d)
}

func main() {
	var myDrink DrinkType = "Pepsi"
	myDrink.selectDrink()
}
