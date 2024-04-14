package printprogramname

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	Run()
	z01.PrintRune('\n')
	// validity: no input or more input
	// if none, print itself
	// if more, print the next one
	// no next directories aka delete /
	// argument = input
	// os.Args = slices of inputs
}

func Run() {
	var originalname string
	var printingname string

	if len(os.Args) > 0 {
		originalname = os.Args[0]
	}

	if len(originalname) > 0 {
		for i := 0; i < len(originalname); i++ {
			if originalname[0] == '/' {
				printingname = string(originalname[i+1])
				if originalname[len(originalname)-1] == '/' {
					printingname = printingname[:len(printingname)-1]
				}
			}
		}
	}

	for _, letter := range printingname {
		z01.PrintRune(letter)
	}
}
