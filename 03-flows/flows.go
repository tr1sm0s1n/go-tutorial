package main

import "fmt"

func main() {
	count := 10
	for i := 0; i < count; i++ {
		if i%2 == 0 {
			fmt.Printf("%d is divisible by 2\n", i)
		} else {
			fmt.Printf("%d isn't divisible by 2\n", i)
			if i == 7 {
				break
			}
		}
	}

	grades := []string{"b", "b", "c", "c", "s", "b", "a"}
	for _, l := range grades {
		switch l {
		case "s":
			fmt.Println("outstanding")
		case "a":
			fmt.Println("excellent")
		default:
			continue
		}
	}
}
