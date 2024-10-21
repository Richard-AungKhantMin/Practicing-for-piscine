package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		return
	}

	fmt.Println("The decimal form of the input is: ", HexToDex(os.Args[1]))

}

func HexToDex(input string) int {

	var IntS []int
	var Modified int

	for i := len(input) - 1; i >= 0; i-- {
		IntS = append(IntS, StrToInt(rune(input[i])))
	}

	for i := 0; i <= len(IntS)-1; i++ {
		Modified = Modified + Power(i)*IntS[i]
	}

	return Modified
}

func StrToInt(c rune) int {

	if c >= '0' && c <= '9' {
		return int(c - '0')
	}

	if c >= 'a' && c <= 'z' {
		return 10 + int(c-'a')
	}

	if c >= 'A' && c <= 'Z' {
		return 10 + int(c-'A')
	}

	return 0
}

func Power(power int) int {

	result := 1

	if power == 0 {
		return result
	}

	for i := 1; i <= power; i++ {
		result = result * 16
	}

	return result
}
