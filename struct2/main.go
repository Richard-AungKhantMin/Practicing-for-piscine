package main

import "fmt"

type meow struct{}

func main() {
	var cat meow

	fmt.Printf("Type of my Struct: %T\n", cat)
	fmt.Println()

	var i int

	fmt.Printf("Type of i: %T\n", i)
}
