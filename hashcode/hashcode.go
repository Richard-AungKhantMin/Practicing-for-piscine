package main

import (
	"fmt"
)

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))

	table := []string{"Z", "Hi!", "BB198365", "sabito", "14 Avril 1912", "zyx987bca", "		pool-2020", "965truma747", " Mercedes-AMG GT"}
	for _, s := range table {
		fmt.Println(HashCode(s))
	}
}

func HashCode(dec string) string {

	size := len(dec)
	hashed := ""

	for _, char := range dec {
		hash := (int(char) + size) % 127
		if hash < 32 || hash > 126 {
			hash += 33
		}
		hashed += string(hash)
	}
	return hashed

}
