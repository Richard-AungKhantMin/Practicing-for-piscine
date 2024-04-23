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

	starter1 := 0
	for i := starter1; i < len(str1)-1; i++ {
		if str1[i] == str1[i+1] && !Idontknow(answer, rune(str1[i])) {
			answer = answer + string(str1[i])
		}
		starter1++
	}

	starter2 := 0
	for i := starter2; i < len(str2)-1; i++ {
		if str2[i] == str2[i+1] && !Idontknow(answer, rune(str2[i])) {
			answer = answer + string(str2[i])
		}
		starter2++
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
