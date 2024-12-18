// This sample program demonstrates how to use a buffered
// channel to work on multiple tasks with a predefined number
// of goroutines.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4  // Number of goroutines to use.
	taskLoad         = 10 // Amount of work to process.
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Create a buffered channel to manage the task load.
	tasks := make(chan string, taskLoad)

	// Launch goroutines to handle the work.
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// Add a bunch of work to get done.
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task: %d", post)
	}

	// Close the channel so the goroutines will quit
	// when all the work is done.
	close(tasks)

	wg.Wait()
}

// worker is launched as a goroutine to process work from
// the buffered channel.
func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		// Wait for work to be assigned.
		task, ok := <-tasks
		if !ok {
			// This means the channel is empty and closed.
			fmt.Printf("Worker %d: Shutting Down\n", worker)
			return
		}

		// Display we are starting the work.
		fmt.Printf("Worker %d: Started %s\n", worker, task)

		// Randomly wait to simulate work time.
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// Display we finished the work.
		fmt.Printf("Worker %d: Completed %s\n", worker, task)
	}
}
