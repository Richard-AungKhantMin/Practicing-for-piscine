package main

import "fmt"

func main() {

	type Person struct {
		name     string
		age      int
		wealth   float64
		virality bool
	}

	var P1 Person = Person{name: "Richard", age: 21, virality: true}
	var P2 Person = Person{name: "Chan Myint", age: 19, wealth: 10000, virality: true}

	people := []Person{P1, P2}

	for _, each := range people {

		fmt.Printf("%s is %d years old. This dude has %v euros in his hand. Shhh! Don't tell anybody. Anyway, it is %t that he's still alive.\n", each.name, each.age, each.wealth, each.virality)

	}
}
