package main

import (
	"fmt"
	"sync"
	"time"
)

// ----------------------
// Define a struct to hold results
// ----------------------

// Result struct represents the output of a worker
// It contains the value (here, image URL) and an error (if any occurred)
type Result struct {
	Value string
	Err   error
}

// ----------------------
// Worker function
// ----------------------

// worker simulates processing an image (or any task) concurrently
// Parameters:
// - url: the task to process (here, image file name)
// - wg: pointer to a WaitGroup to signal when the worker is done
// - resultChan: a channel to send the result back to main
func worker(url string, wg *sync.WaitGroup, resultChan chan Result) {
	defer wg.Done() // tell WaitGroup this worker is done when function exits

	// simulate processing time
	time.Sleep(time.Millisecond * 50)

	// print a message showing which image is processed
	fmt.Printf("Image processed: %s\n", url)

	// send the result to the result channel
	resultChan <- Result{
		Value: url,
		Err:   nil,
	}

	// note: we don’t return anything because results are sent via channel
}

// ----------------------
// Main function
// ----------------------
func main() {
	var wg sync.WaitGroup             // WaitGroup to wait for all goroutines to finish
	resultChan := make(chan Result, 5) // buffered channel to store results (buffer size 5)

	startTime := time.Now() // record start time

	fmt.Println("Welcome to Go Concurrency")

	// Add 4 to the WaitGroup because we will launch 4 workers
	wg.Add(4)

	// Launch 4 worker goroutines
	go worker("Image_1.png", &wg, resultChan)
	go worker("Image_2.png", &wg, resultChan)
	go worker("Image_3.png", &wg, resultChan)
	go worker("Image_4.png", &wg, resultChan)

	// Wait for all workers to finish
	wg.Wait()        // blocks main thread until all Done() are called
	close(resultChan) // close the channel after all workers are done

	// Read results from the channel
	// Using `range` iterates until the channel is closed
	for result := range resultChan {
		fmt.Printf("Received: %v\n", result)
	}

	// Print total elapsed time
	fmt.Printf("It took %s\n", time.Since(startTime))
}

/* ----------------------
Explanation of Concurrency Concepts
---------------------- 

1. Goroutines:
   - Lightweight threads managed by Go runtime.
   - Each call to `go worker(...)` runs concurrently without blocking main.

2. WaitGroup:
   - Used to wait for multiple goroutines to finish.
   - `wg.Add(n)` adds n tasks.
   - Each goroutine calls `defer wg.Done()` when finished.
   - `wg.Wait()` blocks until all tasks are done.

3. Channels:
   - Used for communication between goroutines.
   - `resultChan` carries results from workers to main.
   - Buffered channel (size 5) allows goroutines to send without blocking if space is available.

4. defer wg.Done():
   - Ensures Done() is called even if function exits early (panic or return).

5. Closing channel:
   - `close(resultChan)` is called after all workers finish.
   - Allows `range` to exit when channel is empty.

6. Order of execution:
   - Goroutines may finish in any order.
   - Channels preserve values sent, but the order of reception depends on which goroutine finishes first.

7. Timing:
   - `time.Since(startTime)` shows total time.
   - Because tasks run concurrently, total time is less than sequential execution.
*/
