package main

import "fmt"

type Person struct {
	name  string
	age   int
	alive bool
}

func main() {
	var name string = "Loren" // Variable declaration
	age := 32                 // Type inference
	const pi = 3.14           // Constant declaration
	fmt.Println(name, age, pi)

	var num uint             // Integer
	var deg float32 = 45.5   // Float
	var isTrue bool = true   // Boolean
	var str string = "earth" // String
	fmt.Println(num, deg, isTrue, str)
	fmt.Println(&num, &deg, &isTrue, &str)

	var ptr *uint = &num // Pointer declaration
	num = 7
	fmt.Println(*ptr)

	var p1 Person
	p1.name = "Einstein"
	p1.age = 76
	p1.alive = false
	fmt.Println(p1)

	p2 := Person{"Tyson", 65, true}
	fmt.Println(p2)

	p3 := new(Person)
	fmt.Println(*p3)

	var rolls [3]int // Array declaration
	rolls[2] = 1
	fmt.Println(rolls)

	var planets = []string{"earth", "mars", "venus"} // Slice declaration
	planets = append(planets, "mercury")
	fmt.Println(planets)

	people := make(map[uint]Person) // Map declaration
	people[1] = Person{"Clarke", 90, false}
	people[2] = Person{"Ken", 47, true}
	fmt.Println(people)
}
