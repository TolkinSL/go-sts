package main

import "fmt"

func estimateTime(cups int, seconds int) (totalSeconds int) {
	totalSeconds = cups * seconds
	// return totalSeconds
	return
}

func main() {
	fmt.Println("Total seconds", estimateTime(12, 50))
}