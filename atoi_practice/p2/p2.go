package main

import (
	"fmt"
)

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}

func Atoi(s string) int {

	slicy := []rune(s)
	var answer int
	sign := 1

	if slicy[0] == '+' || slicy[0] == '-' {
		if slicy[0] == '-' {
			sign = -1
		}
		slicy = slicy[1:]
	}

	for _, each := range slicy {
		if each < '0' || each > '9' {
			return 0
		}
	}

	for _, each := range slicy {
		answer = answer*10 + int(each-'0')
	}

	return answer * sign
}
