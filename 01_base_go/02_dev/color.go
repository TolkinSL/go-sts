package main

import (
	"fmt"
	"github.com/fatih/color"
)

func main() {
	fmt.Println("My Color")
	color.Red("Hi Go")

	c := color.New(color.FgCyan).Add(color.Underline)
	c.Println("Prints cyan text.")
	fmt.Printf("%#v\n",c)
}