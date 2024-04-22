package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Read data from standard input
	info, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the data read from standard input
	fmt.Println()
	fmt.Print("Data read from standard input:")
	fmt.Println(string(info))
}
