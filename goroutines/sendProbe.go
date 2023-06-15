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
	go sendProbe(&wg)
	
	wg.Wait()
	fmt.Println("Halting colonization!")
}

func sendProbe(wg *sync.WaitGroup)  {
	defer wg.Done()
	time.Sleep(8*time.Second)
	fmt.Println("Data acquired!")
}