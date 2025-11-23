package main

import "fmt"

func main() {
	dessertMenu := [4]string{"Muff", "Brow", "Croiss", "Cook"}
	fmt.Println(dessertMenu)

	slice := dessertMenu[0:2]
	fmt.Println(slice)

	fullSlice := dessertMenu[:]
	fmt.Println(fullSlice)

	fmt.Printf("%p\n", &dessertMenu[0])
	fmt.Printf("%p\n", &slice[0])
	fmt.Printf("%p\n", &fullSlice[0])

	//--------
	sliceMenu2 := []string{"Muff", "Brow", "Croiss", "Cook"}
	fmt.Println(sliceMenu2)

	slice2 := sliceMenu2[0:2]
	fmt.Println(slice2)

	fullSlice2 := sliceMenu2[:]
	fmt.Println(fullSlice)

	//поле ptr. Его можно увидеть, напечатав адрес первого элемента
	fmt.Printf("%p\n", &sliceMenu2[0])
	fmt.Printf("%p\n", &slice2[0])
	fmt.Printf("%p\n", &fullSlice2[0])

}
