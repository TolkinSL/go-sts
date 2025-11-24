package main

import "fmt"

func deleteByIndex(index int, slice []string) []string {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	drink := []string{"cola", "tea", "water", "pepsi", "coffee"}
	fmt.Printf("%#v\n", drink)
	fmt.Println("len:", len(drink), "cap:", cap(drink))
	fmt.Printf("%p\n", &drink[0])

	fmt.Println("Delete element")
	indexRemove := 2
	drink = append(drink[:indexRemove], drink[indexRemove+1:]...)
	fmt.Printf("%#v\n", drink)
	fmt.Println("len:", len(drink), "cap:", cap(drink))
	fmt.Printf("%p\n", &drink[0])
	fmt.Println()

	fmt.Println("Delete func element")
	drink = deleteByIndex(1, drink)
	fmt.Printf("%#v\n", drink)
	fmt.Println("len:", len(drink), "cap:", cap(drink))
	fmt.Printf("%p\n", &drink[0])
	fmt.Println()

	fmt.Println("for range:")
	menu := []string{"cola", "tea", "water", "pepsi", "coffee"}
	for i, v := range menu {
		fmt.Println(i, v)
	}

	fmt.Println("\nfor i:")
	for i := 0; i < len(menu); i++ {
		fmt.Println(menu[i])
	}
}
