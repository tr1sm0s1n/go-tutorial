package main

import "fmt"

type student struct {
	name, course string
	age, marks   uint
}

func (s student) passed() bool {
	return s.marks >= 50
}

func main() {
	s1 := student{"Eve", "G101", 19, 48}
	s2 := student{"Lam", "G101", 21, 75}
	fmt.Println("Has Eve passed?", s1.passed())
	fmt.Println("Has Lam passed?", s2.passed())
}
