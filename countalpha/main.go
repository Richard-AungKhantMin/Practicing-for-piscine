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

	var count int
	slicy := []rune(s)

	for i := 0; i < len(s); i++ {
		if 'A' <= slicy[i] && slicy[i] <= 'Z' || 'a' <= slicy[i] && slicy[i] <= 'z' {
			count++
		}
	}

	return count
}
