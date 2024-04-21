package main

import (
	"fmt"
)

func main() {
	fmt.Println(CountChar("Hello World", 'l'))
	fmt.Println(CountChar("5  balloons", 5))
	fmt.Println(CountChar("   ", ' '))
	fmt.Println(CountChar("The 7 deadly sins", '7'))
}

func CountChar(str string, c rune) int {
	slicy := []rune(str)
	var count int

	for i := 0; i < len(str); i++ {

		if slicy[i] == c && c == ' ' {
			count++
			break
		}

		if slicy[i] == c {
			count++
		}
	}
	return count
}
