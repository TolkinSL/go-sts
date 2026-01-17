package main

import "fmt"

func main() {
	incomes := []float64{12.3, 32.5, 43.2, 53.2, 11.9, 5.4, 45.1}
	days := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	
	var totalIncome float64

	for i, income := range incomes {
		fmt.Println(days[i], income)
		
		totalIncome += income
	}

	fmt.Println("\nTotal:", totalIncome)

	for _, char := range "Hello" {
		fmt.Printf("%c ", char)
	}

	fmt.Println("")

	drinks := []string{"Cola", "Pepsi", "Tea", "Milk"}

	for _, drink := range drinks {
		if drink == "Tea" {
			continue
		}
		
		fmt.Println(drink)
	}
}
