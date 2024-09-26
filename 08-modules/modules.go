package main

import (
	"fmt"
	"log"

	"go-tutorial/08-modules/helpers"
)

func main() {
	var name string
	var course string
	var marks uint
	fmt.Println("Enter the name:")
	fmt.Scanf("%s", &name)
	fmt.Println("Enter the course:")
	fmt.Scanf("%s", &course)
	fmt.Println("Enter the marks:")
	fmt.Scanf("%d", &marks)

	grade, err := helpers.ConvertMarks(marks)
	if err != nil {
		log.Fatal(err)
	}

	s1 := helpers.Student{Name: name, Course: course, Grade: grade}
	fmt.Println("Here's the report:", s1)
}
