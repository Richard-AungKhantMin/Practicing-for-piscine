package main

import "fmt"

func main() {

	var x int = 42
	var p *int = &x // p now holds the memory address of x

	fmt.Println(&x) // Output: (some memory address, e.g., 0xc0000180a8)
	fmt.Println(p)  // Output: same memory address as &x

}
