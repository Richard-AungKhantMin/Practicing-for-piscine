package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 3 {
		if Check_Gyi(os.Args[1], os.Args[2]) {
			fmt.Println(os.Args[1])
		}
	}
}

func Check_Gyi(first string, second string) bool {

	slicy1 := []rune(first)
	slicy2 := []rune(second)

	final := ""
	var count int = 0

	for i := 0; i < len(slicy1); i++ {
		for j := count; j < len(slicy2); j++ {
			if slicy1[i] == slicy2[j] {
				final = final + string(slicy1[i])
				j = len(slicy2) - 1
			}
			count++
		}
	}

	return final == first
}
