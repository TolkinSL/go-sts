package main

import "fmt"

func getDrink(name string, drink string) {
	fmt.Printf("%s love drink %s\n", name, drink)
}

func loyaltyPoints(amount float64) int {
	points := int(amount * 2)
	return points
}

func updatePoints(current int, newPoints int) int {
	return current + newPoints
}

func main() {
	getDrink("Vasia", "Tea")
	fmt.Println("Points", loyaltyPoints(4.5))

	totalPoints := 120
	totalPoints = updatePoints(totalPoints, loyaltyPoints(4.5))
	fmt.Println("Total Points", totalPoints)
}
