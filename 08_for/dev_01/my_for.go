package main

import "fmt"

func main() {
	cup := 5
	for i := 0; i < cup; i++ {
		fmt.Println("Coffe cup:", i)
	}
	fmt.Println()

	token := 5
	for token > 0 {
		fmt.Println("Token:", token)
		token--
	}
}
