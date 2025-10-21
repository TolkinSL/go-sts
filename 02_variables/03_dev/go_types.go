package main

import "fmt"

func main() {
	price := 4.23
	quantity := 15

	sum := price * float64(quantity)
	fmt.Printf("Summ - %v\n", sum)

	var x, y, z = 1, 1.12, "five"
	fmt.Println(x, y, z)
}