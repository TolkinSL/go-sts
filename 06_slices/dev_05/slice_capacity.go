package main

import "fmt"

func main() {
	myArr := [3]string{"Cola", "Pepsi", "Water"}
	mySlice := myArr[:1]
	fmt.Printf("%#v\n", mySlice)
	fmt.Println("len:", len(mySlice), "cap:", cap(mySlice))
	fmt.Printf("slice: %p\n", &mySlice[0])
	fmt.Println()

	mySlice2 := mySlice

	mySlice = append(mySlice, "Juice")
	fmt.Printf("%#v\n", mySlice)
	fmt.Println("len:", len(mySlice), "cap:", cap(mySlice))
	fmt.Printf("slice: %p\n", &mySlice[0])
	fmt.Println()

	mySlice = append(mySlice, "Tea")
	fmt.Printf("%#v\n", mySlice)
	fmt.Println("len:", len(mySlice), "cap:", cap(mySlice))
	fmt.Printf("slice: %p\n", &mySlice[0])
	fmt.Printf("%#v\n", myArr)
	fmt.Printf("arr: %p\n", &myArr[0])
	fmt.Println()

	mySlice = append(mySlice, "Tequila")
	fmt.Printf("%#v\n", mySlice)
	fmt.Println("len:", len(mySlice), "cap:", cap(mySlice))
	fmt.Printf("slice: %p\n", &mySlice[0])
	fmt.Printf("%#v\n", myArr)
	fmt.Printf("arr: %p\n", &myArr[0])
	fmt.Println()

	mySlice[0] = "Vodka"
	fmt.Printf("%#v\n", mySlice)
	fmt.Println("len:", len(mySlice), "cap:", cap(mySlice))
	fmt.Printf("slice: %p\n", &mySlice[0])
	fmt.Printf("%#v\n", myArr)
	fmt.Printf("arr: %p\n", &myArr[0])
	fmt.Println()

	fmt.Printf("mySlice2 %#v\n", mySlice2)
	fmt.Println("mySlice2 len:", len(mySlice2), "cap:", cap(mySlice2))
	fmt.Printf("mySlice2 slice: %p\n", &mySlice2[0])
	fmt.Println()
}