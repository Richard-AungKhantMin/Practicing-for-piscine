package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 20, 30, 40, 50}
	var length int

	for i := range arr {
		length = i + 1
	}

	fmt.Println("Length of the array is:", length)
}
