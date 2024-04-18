package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))
}

func FromTo(from int, to int) string {

	var answer string

	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid" + "\n"

	}

	if from < to {
		for i := from; i <= to; i++ {
			if i < 10 {
				answer += "0"
			}
			answer += strconv.Itoa(i)
			if i != to {
				answer += ", "
			}
		}
		answer += "\n"
	}

	if from > to {
		for i := from; i >= to; i-- {
			if i < 10 {
				answer += "0"
			}
			answer += strconv.Itoa(i)
			if i != to {
				answer += ", "
			}

		}
		answer += "\n"
	}

	if from == to {
		answer = strconv.Itoa(from) + "\n"
	}
	return answer
}
