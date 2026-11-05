package main

import (
	"fmt"
	"sync"
	"time"
)

// ==========================================
// 1. STRUCT-BASED WORKER EXAMPLE
// ==========================================
type Worker struct {
	ID   int
	Task string
}

func (w *Worker) performTask(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("WorkerID %d started %s\n", w.ID, w.Task)
	time.Sleep(time.Second)
	fmt.Printf("WorkerID %d Finished %s\n", w.ID, w.Task)
}

// ==========================================
// 2. WORKER POOL WITH CHANNELS & WAITGROUP
// ==========================================
func channelWorker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Channel WorkerID %d starting\n", id)

	for task := range tasks {
		time.Sleep(time.Millisecond * 500) // Simulate some work
		results <- task * 2
	}

	fmt.Printf("Channel WorkerID %d finished\n", id)
}

// ==========================================
// 3. SIMPLE WAITGROUP WITHOUT CHANNEL
// ==========================================
func simpleWorker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// wg.Add(1) // Remember: always add before launching goroutine, not inside!
	fmt.Printf("Simple Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Simple Worker %d finished\n", id)
}

func main() {
	// ------------------------------------------
	// RUNNING PATTERN 1: Struct-based Workers
	// ------------------------------------------
	fmt.Println("=== 1. STRUCT WORKER EXAMPLE ===")
	var wg1 sync.WaitGroup

	tasks := []string{"coding", "painting", "driving"}
	for i, task := range tasks {
		work := Worker{
			ID:   i + 1,
			Task: task,
		}
		wg1.Add(1)
		go work.performTask(&wg1)
	}
	wg1.Wait()
	fmt.Println()

	// ------------------------------------------
	// RUNNING PATTERN 2: WaitGroup with Channels
	// ------------------------------------------
	fmt.Println("=== 2. WAITGROUP WITH CHANNELS ===")
	var wg2 sync.WaitGroup
	numWorkers := 3
	numJobs := 3

	tasksChan := make(chan int, numJobs)
	resultsChan := make(chan int, numJobs)

	// Launch workers (Fixed argument order: tasksChan then resultsChan)
	wg2.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go channelWorker(i, tasksChan, resultsChan, &wg2)
	}

	// Send tasks
	for i := 1; i <= numJobs; i++ {
		tasksChan <- i
	}
	close(tasksChan) // Close tasks so workers know when to stop ranging

	// Wait for workers to finish, then close results channel in a separate goroutine
	go func() {
		wg2.Wait()
		close(resultsChan)
	}()

	// Print results
	for result := range resultsChan {
		fmt.Println("Result:", result)
	}
	fmt.Println()

	// ------------------------------------------
	// RUNNING PATTERN 3: Simple WaitGroup
	// ------------------------------------------
	fmt.Println("=== 3. SIMPLE WAITGROUP WITHOUT CHANNEL ===")
	var wg3 sync.WaitGroup
	totalWorkers := 3

	wg3.Add(totalWorkers) // CORRECT WAY: Add counter before launching loop
	for i := 1; i <= totalWorkers; i++ {
		go simpleWorker(i, &wg3)
	}
	wg3.Wait()
	fmt.Println("\nAll simple workers finished!")
}
