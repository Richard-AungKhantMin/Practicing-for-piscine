package main

import (
	"fmt"
)

func main() {

	count := 0
	for burm := 'က'; burm <= 'အ'; burm++ {
		if count%5 == 0 {
			fmt.Println()
		}
		fmt.Printf("The rune number of %c is %d\n", burm, burm)
		count++

	}

}
