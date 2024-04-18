package main

import "fmt"

type Person struct {
	Name            string
	Quote           string
	Years_in_prison int
}

func main() {

	person1 := Person{
		Name:            "Chat Gyi",
		Quote:           "ma pyay nae, ma a loe",
		Years_in_prison: 7,
	}

	fmt.Println("Name:", person1.Name)
	fmt.Println("Quote:", person1.Quote)
	fmt.Println("Years in prison", person1.Years_in_prison)
}
