package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 4
	s := strconv.FormatInt(int64(x), 10)

	fmt.Printf("%#[1]v - %[1]T\n", s)

	var y float64
	fmt.Printf("%#v\n", y)
}