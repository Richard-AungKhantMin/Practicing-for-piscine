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

	fmt.Printf("%s is %d years old. This dude has %v euros in his hand. Shhh! Don't tell anybody. Anyway, it is %t that he's still alive.\n", P1.name, P1.age, P1.wealth, P1.virality)

	fmt.Printf("%s is %d years old. This dude has %v euros in his hand. Shhh! Don't tell anybody. Anyway, it is %t that he's still alive.\n", P2.name, P2.age, P2.wealth, P2.virality)
}
