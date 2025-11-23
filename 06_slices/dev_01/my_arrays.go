package main

import "fmt"

func main() {
	var coffeeSizes [3]string
	fmt.Printf("%#v\n", coffeeSizes)
	
	coffeeSizes[0] = "S"
	fmt.Printf("%#v\n", coffeeSizes)

	coffeeSizes[1] = "L"
	fmt.Printf("%#v\n", coffeeSizes)
}
