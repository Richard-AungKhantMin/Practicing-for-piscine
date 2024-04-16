package piscine

import (
	"fmt"
)

func main() {
	result := ConvertBase("101011", "01", "0123456789")
	fmt.Println(result)
}

func ConvertBase(nbr, baseFrom, baseTo string) string {

	//this is getting the base number. I know this is not a good way of doing it
	//I am being lazy here
	BaseNum1 := len(baseFrom)
	BaseNum2 := len(baseTo)

	//checking what is valid and invalid
	if len(nbr) == 0 || BaseNum1 == 0 || BaseNum2 == 0 {
		return ""
	}

	for i := 0; i < BaseNum1; i++ {
		if baseFrom[i] < '0' || baseFrom[i] > '9' {
			return ""
		}
	}

	for i := 0; i < BaseNum2; i++ {
		if baseTo[i] < '0' || baseTo[i] > '9' {
			return ""
		}
	}

	for i := 0; i < len(nbr); i++ {
		if nbr[i] < '0' || nbr[i] > '9' {
			return ""
		}
	}

	/*
	   Let's change the inputs into int by creating a func atoi

	   After that we get the int values of every input.

	   Create another function that changes nbr in BaseFrom to base 10

	   Create another function that changes the number in base 10 to the number in BaseTo

	   The last function that change the int to string
	*/

}

func myAtoi(a string) int {

	var changed int = 0

	for i := 0; i < len(a); i++ {
		changed = changed*10 + int(rune(a[i])-'0')
	}

	return changed
}
func myItoA(b int) string {

	var changed string = ""
	var slicy []int

	if b > 0 {
		slicy = append(slicy, b%10)
		b = b / 10
	}

	for i := 0; i < len(slicy); i++ {
		changed = changed + string(rune(slicy[i]+'0'))
	}

	return changed
}

func To10(num int) int {

}

func From10(ber int) int {

}
