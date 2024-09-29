package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		return
	}

	fmt.Printf("Arguments: %v\n", os.Args)

	s1 := os.Args[1]
	s2 := os.Args[2]

	if s1 == "" {
		fmt.Println("1")
		return
	}

	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			i++
		}
		j++
	}

	if i == len(s1) {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}
