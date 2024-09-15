package main

import "fmt"

func main() {

	var slicy = []int{1, 3, 5, 7, 9}
	fmt.Println(slicy)

	slicy = append(slicy, 11, 13)

	fmt.Println(slicy)

	var new []int = slicy[3:]
	fmt.Println(new)

}
