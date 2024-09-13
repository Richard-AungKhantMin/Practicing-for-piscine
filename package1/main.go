package main

import (
	"fmt"

	"practicingforpiscine/package1/p1"
	"practicingforpiscine/package1/p2"
)

func main() {

	num1, num2 := 5, 7

	NintMayLin := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(p1.Adder(num1, num2))
	fmt.Println(p2.HalfDaSlice(NintMayLin))

}
