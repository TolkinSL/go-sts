package main

import (
	"fmt"
)

func split(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	left := split(arr[:mid])
	right := split(arr[mid:])

	fmt.Println("left", left, "-", "right", right)
	
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}
	var l, r int

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result
}

func main() {
	arr := []int{7, 3, 6, 2, 8, 11, 15}
	fmt.Println("Отсортированный срез", split(arr))
}
