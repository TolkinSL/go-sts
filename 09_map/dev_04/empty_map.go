package main

import "fmt"

func main() {
	var stock map[string]int
	fmt.Printf("In memory: %p\n", &stock)
	fmt.Println(stock == nil)

	//panic: assignment to entry in nil map
	stock["cola"] = 13
	stock["pespi"] = 23
	fmt.Printf("%+v\n", stock)
}