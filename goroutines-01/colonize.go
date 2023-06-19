package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("Colonizing Mars...")

	wg.Add(1)
	go func() {
		time.Sleep(8 * time.Second)
		fmt.Println("Data acquired!")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Halting colonization!")
}
