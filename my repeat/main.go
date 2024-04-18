package main

import "fmt"

func main() {
	repeat("", 4)
}

func repeat(input string, times int) {
	for i := 1; i <= times; i++ {
		fmt.Println(input)
	}
}
