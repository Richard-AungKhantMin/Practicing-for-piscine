package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	input := os.Args[1]
	var resultS []string
	var word string

	for i := 0; i < len(input); i++ {

		if input[i] != ' ' {
			word = word + string(input[i])
		} else if input[i] == ' ' && i < len(input)-1 && i != 0 && input[i+1] != ' ' {
			resultS = append(resultS, word)
			word = ""
		}

		if i == len(input)-1 {
			//resultS[len(resultS)-1] = resultS[len(resultS)-1] + word
			resultS = append(resultS, word)
		}
	}

	var result string
	for i, eachWord := range resultS {
		if i == 0 {
			result = eachWord
		} else {
			result = result + " " + eachWord
		}
	}

	/*for j, s := range resultS {
		if j < len(resultS)-1 {
			result += s + " "
		} else {
			result += s
		}
	}*/

	PrintStr(result)
}

func PrintStr(word string) {
	for _, each := range word {
		z01.PrintRune(each)
	}
	z01.PrintRune('\n')
}
