package main

import "fmt"

func main() {
	fmt.Println(myAtoi("5090"))
}

func myAtoi(a string) int {

	// It needs validity checking but I have done it in another code which I'm actually coding
	// This is just a testing

	var changed int = 0

	for i := 0; i < len(a); i++ {
		changed = changed*10 + int(rune(a[i])-'0')
	}

	return changed
}
