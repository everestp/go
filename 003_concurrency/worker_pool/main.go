package main

import (
	"fmt"
	"sync"
	"time"
)

// ----------------------
// Define a Result struct
// ----------------------

// Result defines the output structure for processed jobs
// It holds the Value (e.g., image name) and any error (if occurred)
type Result struct {
	Value string
	Err   error
}

// ----------------------
// Worker function
// ----------------------

// worker takes jobs from jobsChan, processes them, and sends output to resultChan
func worker(jobsChan chan string, wg *sync.WaitGroup, resultChan chan Result) {
	defer wg.Done() // Decrement WaitGroup counter when this worker finishes

	// Loop over jobs until jobsChan is closed
	for job := range jobsChan {
		// Simulate processing time (here, 5 seconds per image)
		time.Sleep(time.Millisecond * 5000)

		// Print message to show which job is processed
		fmt.Printf("Image processed: %s\n", job)

		// Send result back to the results channel
		resultChan <- Result{
			Value: job,
			Err:   nil,
		}
	}
}

// ----------------------
// Main function
// ----------------------
func main() {
	// List of images to process
	jobs := []string{
		"image-1.png", "image-2.png", "image-3.png", "image-4.png", "image-5.png",
		"image-6.png", "image-7.png", "image-8.png", "image-9.png", "image-10.png",
		"image-11.png", "image-12.png", "image-13.png", "image-14.png", "image-15.png",
		"image-16.png", "image-17.png", "image-18.png", "image-19.png", "image-20.png",
	}

	var wg sync.WaitGroup          // WaitGroup to wait for all workers
	totalWorkers := 5             // Number of concurrent workers

	// Channels for communication
	resultChain := make(chan Result, len(jobs)) // Buffered to avoid blocking
	jobsChan := make(chan string, len(jobs))    // Buffered to hold all jobs

	startTime := time.Now() // Record start time

	// ----------------------
	// 1. Initialize worker goroutines
	// ----------------------
	for i := 1; i <= totalWorkers; i++ {
		wg.Add(1) // Increment WaitGroup counter
		go worker(jobsChan, &wg, resultChain)
	}

	// ----------------------
	// 2. Send jobs to jobsChan
	// ----------------------
	for i := 0; i < len(jobs); i++ {
		jobsChan <- jobs[i] // Send each job into the channel
	}

	// Close jobsChan to signal workers no more jobs are coming
	close(jobsChan)

	// ----------------------
	// 3. Monitor completion in background
	// ----------------------
	go func() {
		wg.Wait()          // Wait for all workers to finish
		close(resultChain) // Close results channel so range can exit
	}()

	// ----------------------
	// 4. Collect and print results
	// ----------------------
	for result := range resultChain {
		fmt.Printf("Received: %v\n", result.Value)
	}

	fmt.Printf("Time taken := %v\n", time.Since(startTime))
}

/* ----------------------
Detailed Explanation of Concurrency Concepts
----------------------

1. Worker Pool:
   - We have multiple workers (goroutines) processing jobs concurrently.
   - totalWorkers = 5 means 5 images can be processed at the same time.

2. Channels:
   - jobsChan: holds jobs to be processed by workers.
   - resultChain: collects results from workers.
   - Buffered channels prevent workers from blocking if main is not reading immediately.

3. WaitGroup:
   - Keeps track of active workers.
   - wg.Add(1) for each worker, wg.Done() when finished.
   - wg.Wait() blocks until all workers finish.

4. Closing Channels:
   - Close jobsChan after sending all jobs to signal "no more work".
   - Close resultChain after all workers finish so range loop can exit.

5. Goroutines:
   - Each worker is a goroutine running concurrently.
   - The anonymous goroutine for wg.Wait() ensures resultChain closes when all done.

6. Advantages:
   - Efficient concurrency: multiple images processed simultaneously.
   - Clean code: workers are decoupled from main.
   - Easy to scale: increase totalWorkers to process more jobs concurrently.

7. Output Order:
   - Since workers run concurrently, results may arrive in **different order** than jobs list.
   - Channels preserve the data sent, but the order depends on which worker finishes first.
*/
