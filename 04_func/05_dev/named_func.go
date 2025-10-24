package main

import "fmt"

func main() {
	taxRate := 0.10

	calcTax := func(amount float64) float64 {
		return amount * taxRate
	}

	var subtotal float64 = 15
	total := subtotal + calcTax(subtotal)
	fmt.Println("Total sub + tax $", total)
}
