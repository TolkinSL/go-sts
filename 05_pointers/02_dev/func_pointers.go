package main

import "fmt"

func calcDiscountPrice(price float64, discount float64) float64 {
	fmt.Println("Memory location price", &price)
	return price - price * discount
}

func applyDiscount(price *float64, discount float64) {
	*price = *price - (*price * discount)
}

func main() {
	var coffePrice float64 = 5.67
	fmt.Println("Base coffe price", coffePrice)
	fmt.Println("Memory location coffePrice", &coffePrice)
	fmt.Println("Discount price", calcDiscountPrice(coffePrice, 0.1))

	fmt.Println("Base coffe price", coffePrice)
	applyDiscount(&coffePrice, 0.1)
	fmt.Println("Base coffe price change from pointers", coffePrice)
}