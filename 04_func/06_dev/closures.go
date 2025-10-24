package main

import "fmt"

func createRegulator() func(change float64) float64 {
	baseTemp := 90.0

	adjustTemp := func(change float64) float64 {
		baseTemp += change
		fmt.Println("adjustTemp:", baseTemp)
		return baseTemp
	}

	adjustTemp(1)
	fmt.Println("createRegulator:", baseTemp)

	adjustTemp(-2)
	fmt.Println("createRegulator:", baseTemp)

	return adjustTemp

}

func main() {
	test := createRegulator()
	fmt.Println("test +5", test(5))	
	fmt.Println("test -3", test(-3))
}
