package main

import (
	"fmt"
)

func main() {
	table := [][]string{
		{"abc", "def"},
		{"hello", "yoall"},
		{"everyone", ""},
		{"hello world", "fam"},
		{"abc", "abc"},
		{"", ""},
		{"pomme", "pomme"},
		{"+265", "265"},
		{"123231", "123231"},
		{"w^p@@j", "w^p@@j"},
		{"26235e5", "4478q92"},
		{"		", "		 "},
		{"AB$%d.52", "eepqdl.52"},
		{"", "eveRyone"},
		{"_55w1se", "55w1se"},
	}
	for _, arg := range table {
		fmt.Println(WeAreUnique(arg[0], arg[1]))
	}
}

func WeAreUnique(str1, str2 string) (int, string) {

	var answer string

	if len(str1) == 0 && len(str2) == 0 {
		return -1, ""
	}

	if len(str1) == 0 {
		for _, each := range str2 {
			if !IsContain(str2, each) && !IsContain(answer, each) {
				answer = answer + string(each)
			}
		}
	}

	if len(str2) == 0 {
		for _, each := range str1 {
			if !IsContain(str1, each) && !IsContain(answer, each) {
				answer = answer + string(each)
			}
		}
	}

	for _, each := range str1 {
		if !IsContain(str2, each) && !IsContain(answer, each) {
			answer = answer + string(each)
		}
	}

	for _, each := range str2 {
		if !IsContain(str1, each) && !IsContain(answer, each) {
			answer = answer + string(each)
		}
	}

	return len(answer), answer
}

func IsContain(word string, char rune) bool {

	for _, each := range word {
		if each == char {
			return true
		}
	}
	return false
}
