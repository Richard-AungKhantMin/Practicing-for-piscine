package main

import "fmt"

func main() {

	var answer string

	example := "WhatDaDogDoing?"
	word := "Do Do Do Doing!"

	for _, each := range word {
		if containRune(example, each) && !containRune(answer, each) {
			answer = answer + string(each)
		}
	}

	fmt.Println(answer)

}

// This approach is for my dumb brain. When nested loops are confusing enough, make it into another funtion.

func containRune(woof string, meow rune) bool {
	for _, each := range woof {
		if each == meow {
			return true
		}
	}
	return false
}