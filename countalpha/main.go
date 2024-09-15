package main

import (
	"fmt"
)

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}

func CountAlpha(s string) int {

	count := 0

	for _, each := range s {

		if each >= 'a' && each <= 'z' || each >= 'A' && each <= 'Z' {
			count++
		}

	}

	return count
}
