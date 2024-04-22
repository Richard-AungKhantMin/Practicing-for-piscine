package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if the correct number of arguments are passed
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./quadC <arg1> <arg2>")
		os.Exit(1)
	}

	// Convert arguments from strings to integers
	arg1, err1 := strconv.Atoi(os.Args[1])
	arg2, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil {
		fmt.Println("Both arguments must be integers.")
		os.Exit(1)
	}

	// Example function that takes two integers
	result := exampleFunction(arg1, arg2)

	// Output the result
	fmt.Println("Result:", result)
}

// exampleFunction performs a sample calculation or any specific task
func exampleFunction(a, b int) int {
	// Placeholder logic
	return a + b // Just an example, replace with actual logic
}
