package main

import (
	"github.com/01-edu/z01"
)

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintMemory(arr [10]byte) {
	// Print hexadecimal memory representation
	for i, c := range arr {
		hexDigit1 := c / 16 // Get the first 4 bits (high bits)
		hexDigit2 := c % 16 // Get the last 4 bits (low bits)

		// Convert each hexadecimal digit to ASCII representation and print
		z01.PrintRune(hexDigitToRune(hexDigit1))
		z01.PrintRune(hexDigitToRune(hexDigit2))

		z01.PrintRune(' ')

		// Go to another line after printing 4 items
		if i%4 == 3 {
			z01.PrintRune('\n')
		}
	}
	z01.PrintRune('\n')

	// Print ASCII characters
	for _, char := range arr {
		// Check if char is printable
		if char >= 32 && char <= 126 {
			z01.PrintRune(rune(char))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}

func hexDigitToRune(digit byte) rune {
	if digit < 10 {
		return rune('0' + digit)
	}
	return rune('a' + (digit - 10))
}
