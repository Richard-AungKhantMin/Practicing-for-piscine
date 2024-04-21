package main

import (
	"fmt"
)

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}

func StrLen(s string) int {
	var count int
	tempo := []rune(s)

	for i := 0; i < len(tempo); i++ {
		count++
	}

	return count
}
