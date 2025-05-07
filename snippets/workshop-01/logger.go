package main

import (
	// "context"
	"log"
	"snippets/workshop-01/types"
	"time"
)

// printTasksPeriodicaly demonstrates channels and goroutines
func logTasksPeriodically(data *types.Data, done chan bool) {
	ticker := time.NewTicker(10 * time.Second)
	count := 0

	for {
		select {
		case <-ticker.C:
			log.Println("Current tasks:")

			// Example of flow control with break and continue
			for id, task := range data.Tasks {
				// Skip completed tasks as example of continue
				if task.Completed {
					continue
				}

				log.Printf("ID: %d, Title: %s\n", id, task.Title)
				count++

				// Break after printing 5 tasks as example of break
				if count >= 5 {
					log.Println("Reached maximum tasks to display")
					break
				}
			}
			count = 0

		case <-done:
			log.Println("Task monitoring stopped")
			return
		}
	}
}
