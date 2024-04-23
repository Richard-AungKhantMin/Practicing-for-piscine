package main

import (
	"fmt"
)

func main() {
	fmt.Println(Itoa(-2))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}

func Itoa(n int) string {

	var answer string
	var slicy []int

	if n < 0 {
		answer = "-"
		n = n * -1
	}

	if n == 0 {
		return "0"
	}

	for n > 0 {
		slicy = append(slicy, n%10)
		n = n / 10
	}

	for i := len(slicy) - 1; i >= 0; i-- {

		answer = answer + string(rune(slicy[i])+'0')
	}

	return answer
}
