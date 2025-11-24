package main

import "fmt"

func main() {
	menu := []string{"Tea", "Coffee"}
	fmt.Printf("%#v\n", menu)
	fmt.Println("Len:", len(menu), "Cap:", cap(menu))
	fmt.Println()

	_ = append(menu, "Cola")
	fmt.Printf("%#v\n", menu)
	fmt.Printf("Переменная: %p\n", &menu)
	fmt.Printf("Указатель: %p\n", &menu[0])
	fmt.Println()

	menu = append(menu, "Cola")
	fmt.Printf("Переменная: %p\n", &menu)
	fmt.Printf("Указатель: %p\n", &menu[0])
	fmt.Println("Len:", len(menu), "Cap:", cap(menu))
	fmt.Println()

	menu = append(menu, "Water")
	fmt.Printf("Переменная: %p\n", &menu)
	fmt.Printf("Указатель: %p\n", &menu[0])
	fmt.Println("Len:", len(menu), "Cap:", cap(menu))
	fmt.Println()

	menu = append(menu, "Milk")
	fmt.Printf("Переменная: %p\n", &menu)
	fmt.Printf("Указатель: %p\n", &menu[0])
	fmt.Println("Len:", len(menu), "Cap:", cap(menu))
	fmt.Println()

	testCap := make([]int, 7)
	fmt.Printf("%#v\n", testCap)
	fmt.Println("len:", len(testCap), "cap:", cap(testCap))
	fmt.Println()

	testCap[0] = 2
	fmt.Println("len:", len(testCap), "cap:", cap(testCap))
	fmt.Printf("%#v\n", testCap)
	fmt.Println()

	testCap = append(testCap, 4)
	fmt.Println("len:", len(testCap), "cap:", cap(testCap))
	fmt.Printf("%#v\n", testCap)
	fmt.Println()

	testCap2 := make([]int, 0)
	fmt.Printf("%#v\n", testCap2)
	fmt.Println("len:", len(testCap2), "cap:", cap(testCap2))
	fmt.Println()
	// testCap2[0] = 2 runtime error: index out of range [0]
}
