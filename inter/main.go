package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		return
	}

	first := os.Args[1]
	second := os.Args[2]
	var answer string

	for _, each := range first {
		if Finder(second, each) && !Finder(answer, each) {
			answer = answer + string(each)
		}
	}
	fmt.Println(answer)
}

func Finder(Big string, small rune) bool {

	for _, each := range Big {
		if each == small {
			return true
		}
	}
	return false
}
