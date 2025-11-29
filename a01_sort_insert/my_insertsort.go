package main

import "fmt"

func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; {
			arr[j-1], arr[j] = arr[j], arr[j-1]
			j -= 1
		}
	}
}

func main() {
	arr := []int{2, 3, 4, 1, 6}
	fmt.Println("Первоначальный массив")
	fmt.Println(arr)

	insertSort(arr)
	fmt.Println("Отсортированный массив")
	fmt.Println(arr)
}
