package main

import (
	"fmt"
	"sync"
	"time"
)

// Result defines the output structure for our processed jobs
type Result struct {
	Value string
	Err   error
}

// worker takes jobs from jobsChan, processes them, and sends output to resultChan
func worker(jobsChan chan string, wg *sync.WaitGroup, resultChan chan Result) {
	defer wg.Done() // Ensure the counter decrements when the worker finishes

	for job := range jobsChan {
		// Simulate processing time
		time.Sleep(time.Millisecond * 5000)
		
		fmt.Printf("image processed: %s\n", job)

		// Send result back to the results channel
		resultChan <- Result{
			Value: job,
			Err:   nil,
		}
	}
}

func main() {
	// The list of images to process
	jobs := []string{
		"image-1.png", "image-2.png", "image-3.png", "image-4.png", "image-5.png",
		"image-6.png", "image-7.png", "image-8.png", "image-9.png", "image-10.png",
		"image-11.png", "image-12.png", "image-13.png", "image-14.png", "image-15.png",
		"image-16.png", "image-17.png", "image-18.png", "image-19.png", "image-20.png",
	}

	var wg sync.WaitGroup
	totalWorkers := 5
	
	// Channels for communication
	// resultChain is buffered to prevent workers from blocking
	resultChain := make(chan Result, len(jobs))
	jobsChan := make(chan string, len(jobs))
	
	startTime := time.Now()

	// 1. Initialize Workers
	for i := 1; i <= totalWorkers; i++ {
		wg.Add(1)
		go worker(jobsChan, &wg, resultChain)
	}

	// 2. Feed the jobs into the channel
	for i := 0; i < len(jobs); i++ {
		jobsChan <- jobs[i]
	}
	
	// Close jobsChan so workers know there is no more work coming
	close(jobsChan)

	// 3. Monitor for completion in the background
	go func() {
		wg.Wait()          // Wait for all workers to finish their loop
		close(resultChain) // Close the results so the range loop below can end
	}()

	// 4. Collect and print results
	for result := range resultChain {
		fmt.Printf("Received : %v\n", result.Value)
	}

	fmt.Printf("Time taken := %v\n", time.Since(startTime))
}