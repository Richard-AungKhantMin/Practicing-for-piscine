package main

import (
	"fmt"
)

func main() {
	QuadA(5, 3)
	QuadA(1, 0)
	QuadA(8, 0)
	QuadA(1, 1)
	QuadA(-1, 9)
	QuadA(5, -2)
}

func QuadA(x, y int) string {
	if x < 0 || y < 0 {
		fmt.Println("Invalid input, bro. Are you crazy?")
	}

	if x == 1 && y > 1 {
		for i := 1; i <= y; i++ {
			if i == 1 {
				fmt.Println("0")
			}

			if 2 <= x && i < y {
				fmt.Println("|")
			}

			if i == y {
				fmt.Println("o")
			}
		}
	}

	if x > 1 && y == 0 {

	}

	if x == 0 && y == 0 {
		fmt.Println("o")
	}

	if x > 0 && y > 0 {

	}

}
