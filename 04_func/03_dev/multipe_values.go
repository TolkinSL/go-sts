package main

import "fmt"

func processPayment(orderTotal, tip, paid float64) (float64, float64) {
	totalAmount := orderTotal + tip
	change := paid - totalAmount
	return totalAmount, change
}

func main() {
	totalCost, change := processPayment(6.50, 2.0, 10)
	fmt.Println(totalCost, change)
}
