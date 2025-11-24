package main

import "fmt"

func main() {
	menu1 := [4]string{"Tea", "Juice", "Coffee", "Cola"}
	slices1 := menu1[:3]

	fmt.Printf("%#v\n", slices1)
	fmt.Printf("arr			%p\n", &menu1[0])
	fmt.Printf("slices	%p\n", &slices1[0])

	slices1[1] = "Pepsi"
	fmt.Printf("%#v\n", slices1)
	fmt.Printf("%#v\n", menu1)

	slices2 := []int{2, 5, 8}
	fmt.Printf("%+v\n", slices2)

	drinkTypes := make([]string, 3)
	drinkTypes[1] = "Pepsi"
	fmt.Printf("%#v\n", drinkTypes)
}
