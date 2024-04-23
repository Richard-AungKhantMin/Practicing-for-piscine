package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println()
	}

	var answer string
	first := os.Args[1]
	second := os.Args[2]

	for _, each1 := range first {
		if !IsThatHere(answer, each1) {
			answer = answer + string(each1)
		}
	}

	for _, each2 := range second {
		if !IsThatHere(answer, each2) {
			answer = answer + string(each2)
		}
	}
	fmt.Println(answer)
}

func IsThatHere(Inside string, DaWord rune) bool {
	for _, each := range Inside {
		if each == DaWord {
			return true
		}
	}
	return false
}
