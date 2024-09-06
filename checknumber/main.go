package main

import (
	"fmt"
)

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}

func CheckNumber(arg string) bool {

	for _, each := range arg {
		if '0' <= each && each <= '9' {
			return true
		}
	}
	return false
}
