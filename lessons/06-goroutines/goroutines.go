package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("Hello from Earth")

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Hello from Mars")
	}()

	fmt.Println("Hello again from Earth")
	wg.Wait()
}
