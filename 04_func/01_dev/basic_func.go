package main

import "fmt"

func greet() {
	fmt.Println("Hello")
}

func greetName(name string) {
	fmt.Println("Hello", name)
}

func main() {
	greet()
	greetName("Vasia")
}
