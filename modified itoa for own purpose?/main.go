package main

import "fmt"

func main() {
	fmt.Println(691005)
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
