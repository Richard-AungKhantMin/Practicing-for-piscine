package main

import (
	"fmt"
)

func main() {
	fmt.Println(Gcd(42, 10))
	fmt.Println(Gcd(42, 12))
	fmt.Println(Gcd(14, 77))
	fmt.Println(Gcd(17, 3))

	fmt.Println(Gcd(12, 23))
	fmt.Println(Gcd(25, 15))
	fmt.Println(Gcd(23043, 122))
	fmt.Println(Gcd(11, 77))
	fmt.Println(Gcd(15, 30))
}

func Gcd(a, b uint) (uint, []uint, []uint) {

	Af := factor(a)
	Bf := factor(b)
	var answer uint

	for _, each1 := range Af {
		for _, each2 := range Bf {

			if each1 == each2 {
				answer = each1
			}

		}
	}

	return answer, Af, Bf
}

func factor(a uint) []uint {

	var answer []uint
	for i := uint(1); i <= a; i++ {

		if a%i == 0 {

			answer = append(answer, i)

		}

	}

	return answer
}

/*

Find the factors
Get the same factors
Compare them for the biggest one

*/
