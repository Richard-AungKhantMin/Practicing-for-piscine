package main

import (
	"fmt"

	// fix the module address

	"practicingforpiscine/package1/p1"
	"practicingforpiscine/package1/p2"
)

func main() {

	num1, num2 := 5, 7

	NintMayLin := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	NintMayLin[2] = 1978
	NintMayLin[3] = 556

	fmt.Println(p1.Adder(num1, num2))
	fmt.Println(p2.HalfDaSlice(NintMayLin))

}
