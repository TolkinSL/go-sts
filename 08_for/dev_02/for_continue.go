package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Terminal")

	count := 0

	var coffeeName string
	for {
		fmt.Print("Enter cofee: ")
		_, err := fmt.Scanln(&coffeeName)
		fmt.Println("My cofee:", coffeeName)
		count++

		if coffeeName == "exit" {
			break
		}

		if err != nil {
			fmt.Println("Enter valid cofee: ")
			continue
		}

	}
	
	fmt.Println(count, "Finish order...")
}