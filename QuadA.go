package main

import (
	"fmt"
)

func main() {

	/*
		QuadA(5, 3)
		QuadA(1, 5)

		QuadA(1, 0)
		QuadA(0, 1)
		QuadA(8, 0)
		QuadA(0, 13)

		QuadA(1, 1)

		QuadA(-1, 9)

		QuadA(9, 1)

		QuadA(5, -2)
	*/

	QuadA(10, 10)
}

func QuadA(x, y int) {
	//invalid
	if x <= 0 || y <= 0 {
		fmt.Println()
		return
	}

	// for the case, (1, not 1)
	if x == 1 && y > 1 {
		for i := 1; i <= y; i++ { //creating rooms according to the input
			if i == 1 {
				fmt.Println("o")
			}

			if 2 <= y && i < y { //I used i < y instead of i= < y-1. It's the same.
				fmt.Println("|")
			}

			if i == y {
				fmt.Println("o")
			}
		}

		return
	}

	// for the case (not 1, 1)
	if x > 1 && y == 1 {
		for i := 1; i <= x; i++ {

			if i == 1 {
				fmt.Printf("%s", "o")
			}

			if 2 <= i && i < x {
				fmt.Printf("%s", "-")
			}

			if i == x {
				fmt.Printf("%s", "o")
			}
		}
		fmt.Println()
		return
	}

	// (1,1) case
	if x == 1 && y == 1 {
		fmt.Println("o")
		return
	}

	// a very typical square
	if x > 0 && y > 0 {

		//1st line
		for i := 1; i <= x; i++ {

			if i == 1 {
				fmt.Printf("%s", "o")
			}

			if 2 <= i && i < x {
				fmt.Printf("%s", "-")
			}

			if i == x {
				fmt.Printf("%s", "o")
			}
		}
		fmt.Println() // printing next line after 1st horizontal line

		//the middle lines
		for i := 1; i < y-1; i++ { // rows no. 2 to the one before the last
			fmt.Printf("%s", "|") // first | string

			for j := 2; j <= x-1; j++ {
				fmt.Printf("%s", " ")
			}

			fmt.Printf("%s", "|") // last | string
			fmt.Println()
		}

		// the last line
		for i := 1; i <= x; i++ {

			if i == 1 {
				fmt.Printf("%s", "o")
			}

			if 2 <= i && i < x {
				fmt.Printf("%s", "-")
			}

			if i == x {
				fmt.Printf("%s", "o")
			}

		}
		fmt.Println()
		return
	}

}
