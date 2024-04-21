package main

import (
	"fmt"
)

func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}

func DigitLen(n, base int) int {

	var count int

	if n < 0 {
		n = n * -1
	}

	if base < 0 {
		return -1
	}

	for n > 0 {
		n = n / base
		count++
	}

	return count
}
