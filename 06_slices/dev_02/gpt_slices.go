package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func hdrData(s []int) uintptr {
	return (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data
}

func main() {
	s0 := []int{10, 20, 30}
	s1 := s0[1:3]
	s2 := s0[:]

	fmt.Printf("&s0 header: %p\n", &s0)
	fmt.Printf("&s1 header: %p\n", &s1)
	fmt.Printf("&s2 header: %p\n\n", &s2)

	fmt.Printf("s0.data (ptr): 0x%x\n", hdrData(s0))
	fmt.Printf("s1.data (ptr): 0x%x\n", hdrData(s1))
	fmt.Printf("s2.data (ptr): 0x%x\n\n", hdrData(s2))

	fmt.Printf("&s0[0]: %p\n", &s0[0])
	fmt.Printf("&s1[0]: %p (== &s0[1])\n", &s1[0])
	fmt.Printf("&s2[0]: %p\n\n", &s2[0])

	// Демонстрация append (возможно перераспределение)
	oldData := hdrData(s0)
	s0 = append(s0, 40, 50, 60) // может перераспределить
	fmt.Printf("after append, s0.data: 0x%x (old 0x%x)\n", hdrData(s0), oldData)
}
