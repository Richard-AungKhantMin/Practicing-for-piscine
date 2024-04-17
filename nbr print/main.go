package main

import (
	"github.com/01-edu/z01"
)

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}

func PrintNbr(n int) {
	var sign rune
	var newV int
	var collector []rune

	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		sign = '-'
		newV = n * -1
	} else {
		newV = n
	}

	for newV > 0 {
		collector = append(collector, rune((newV%10)+'0'))
		newV = newV / 10
	}

	if n < 0 {
		z01.PrintRune(sign)
		for i := len(collector) - 1; i >= 0; i-- {
			z01.PrintRune(collector[i])
		}
		return
	}

	if n > 0 {
		for i := len(collector) - 1; i >= 0; i-- {
			z01.PrintRune(collector[i])
		}
		return
	}
}
