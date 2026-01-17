package main

import "fmt"

func main() {
	var stock map[string]int
	fmt.Printf("%#v\n", stock)

	stock = make(map[string]int)
	fmt.Printf("%#v\n", stock)

	stock["cola"] = 14
	fmt.Printf("%#v\n", stock)
}