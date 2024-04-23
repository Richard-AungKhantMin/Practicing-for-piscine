package main

import (
	"fmt"
)

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}

func WeAreUnique(str1, str2 string) int {

	var answer string

	if len(str1) == 0 || len(str2) == 0 {
		return -1
	}

	for _, each := range str1 {
		if !Idontknow(str2, each) && !Idontknow(answer, each) {
			answer = answer + string(each)
		}
	}

	for _, each2 := range str2 {
		if !Idontknow(str1, each2) && !Idontknow(answer, each2) {
			answer = answer + string(each2)
		}
	}

	return len(answer)
}

func Idontknow(checkgyi string, checklay rune) bool {

	for _, sautkalay := range checkgyi {
		if sautkalay == checklay {
			return true
		}
	}
	return false
}
