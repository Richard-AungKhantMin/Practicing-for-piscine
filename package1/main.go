package main

import (
	"fmt"

	"practicingforpiscine/package1/p1"
	"practicingforpiscine/package1/p2"
)

func main() {

	num1, num2 := 5, 7

	NintMayLin := []int{1, 3, 5, 7, 9}

	fmt.Println(p1.Adder(num1, num2))
	fmt.Println(p2.HalfDaSlice(NintMayLin))

}
