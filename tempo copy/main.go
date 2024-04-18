package main

import (
	"github.com/01-edu/z01"
)

func main() {
	answer := "x = 42, y = 21"
	for i := 0; i < len(answer); i++ {
		z01.PrintRune(rune(answer[i]))
	}
	z01.PrintRune('\n')
}
