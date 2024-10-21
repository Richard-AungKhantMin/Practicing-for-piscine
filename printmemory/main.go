package main

import (
	"github.com/01-edu/z01"
)

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})

}

func PrintMemory(arr [10]byte) {

	//First part
	for i, c := range arr {
		hexDigit1 := c / 16
		hexDigit2 := c % 16

		z01.PrintRune(hexDigitToRune(hexDigit1))
		z01.PrintRune(hexDigitToRune(hexDigit2))

		z01.PrintRune(' ')

		if i%4 == 3 {
			z01.PrintRune('\n')
		}
	}
	z01.PrintRune('\n')

	//Just printing
	for _, char := range arr {

		if char >= 32 && char <= 126 {
			z01.PrintRune(rune(char))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}

// changing to rune
func hexDigitToRune(digit byte) rune {
	if digit < 10 {
		return rune('0' + digit)
	}
	return rune('a' + (digit - 10))
}
