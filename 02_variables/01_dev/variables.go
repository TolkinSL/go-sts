package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Medium Espresso")

	var coffeName string = "Espresso"
	var size = "Small"
	var price = 2.50
	count := 3
	fmt.Println(size, coffeName, "price $", price)
	fmt.Printf("%s %s price $%.2f\n", size, coffeName, price)

	fmt.Printf("Type of size - %T\n", size)
	fmt.Println(unsafe.Sizeof(count))
}
