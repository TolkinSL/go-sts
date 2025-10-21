package main

import "fmt"

func main() {

	// Untyped constant
	const rewardPoints = 10
	fmt.Printf("reward type - %T\n", rewardPoints)

	var totalPoints float64 = 134.6
	totalPoints += rewardPoints
}
