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

	person2 := Person{
		Name:            "Aung Khant Min",
		Quote:           "Sate sin yell dl",
		Years_in_prison: 0,
	}

	person3 := Person{
		Name:            "Ni hao ma",
		Quote:           "bing chilling",
		Years_in_prison: 921473927,
	}

	fmt.Println("Name:", person1.Name)
	fmt.Println("Quote:", person1.Quote)
	fmt.Println("Years in prison", person1.Years_in_prison)
	fmt.Println()

	fmt.Println(person2.Name)
	fmt.Println(person2.Quote)
	fmt.Println(person2.Years_in_prison)
	fmt.Println()

	fmt.Println("Name-", person3.Name)
	fmt.Println("Quote-", person3.Quote)
	fmt.Println("Years in prison -", person3.Years_in_prison)
}
