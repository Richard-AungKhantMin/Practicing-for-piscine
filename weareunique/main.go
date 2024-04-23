package main

import (
	"fmt"
)

func main() {
	fmt.Println(WeAreUnique("foo", "boobb"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}

func WeAreUnique(str1, str2 string) int {

	var answer string
	var tempo string
	combined := str1 + str2

	if len(str1) == 0 || len(str2) == 0 {
		return -1
	}

	for _, each := range str1 {
		if Idontknow(str2, each) && !Idontknow(answer, each) {
			answer = answer + string(each)
		}
	}

	for _, each := range str2 {
		if Idontknow(str1, each) && !Idontknow(answer, each) {
			answer = answer + string(each)
		}
	}

	for i := 0; i < len(str1); i++ {
		for j := i + 1; j < len(str1); j++ {
			if str1[i] == str1[j] && !Idontknow(answer, rune(str1[i])) {
				answer = answer + string(str1[i])
			}
		}
	}

	for i := 0; i < len(str2); i++ {
		for j := i + 1; j < len(str2); j++ {
			if str2[i] == str2[j] && !Idontknow(answer, rune(str2[i])) {
				answer = answer + string(str2[i])
			}
		}
	}

	for _, each := range combined {
		if !Idontknow(answer, each) {
			tempo = tempo + string(each)
		}
	}

	return len(tempo)
}

func Idontknow(checkgyi string, checklay rune) bool {

	for _, sautkalay := range checkgyi {
		if sautkalay == checklay {
			return true
		}
	}
	return false
}
