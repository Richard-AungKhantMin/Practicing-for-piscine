package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 1, 2, 3, 4}
	unmatch := Unmatch(a)
	fmt.Println(unmatch)
}

func Unmatch(a []int) int {

	var count int

	for i := 0; i < len(a); i++ {
		count = 0
		for j := 0; j < len(a); j++ {
			if a[i] == a[j] {
				count++
			}
		}
		if count%2 != 0 {
			return a[i]
		}
	}

	return -1
}
