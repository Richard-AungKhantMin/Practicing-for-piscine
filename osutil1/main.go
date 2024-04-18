package main

import (
	"fmt"
	// "io/ioutil" is no longer in Go 1.16
	"os"
)

func main() {

	filename, err := os.ReadFile("example.txt")
	//stored as a byte. os.ReadFile usage

	if err != nil {
		fmt.Printf("The error was: %v\n", err)
		return
	}

	fmt.Println(string(filename))

}
