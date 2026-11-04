package main

import (
	"fmt"
	"sync"
	"time"
)

//CONSTRCUTION EXAMPLE

type Worker struct {
	ID   int
	Task string
}

func (w *Worker) performTask(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("WorderID %d  started %s \n", w.ID, w.Task)
	time.Sleep(time.Second)
	fmt.Printf("WorderID %d  Finished %s \n", w.ID, w.Task)
}

func main() {

	var wg sync.WaitGroup

	//Define  task top be performed by workers
	tasks := []string{"coding", "painting", "driving"}
	for i, task := range tasks {

		work := Worker{
			ID:   i,
			Task: task,
		}
		wg.Add(1)
		go work.performTask(&wg)
	}
	// Wait for all workers to finish
	wg.Wait()
}

///======= WAITGROUP  WITH CHANNEL========

// func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("WorkerID %d starting \n", id)
// 	time.Sleep(time.Second) //Simulate some work
// 	for task := range tasks {

// 		results <- task * 2
// 	}
// 	fmt.Printf("WordkerID %d finished \n", id)

// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3
// 	numsJobs := 3

// 	results := make(chan int, numsJobs)
// 	tasks := make(chan int, numsJobs)
// 	wg.Add(numWorkers)

// 	for i := range numWorkers {
// 		go worker(i, results, tasks, &wg)
// 		go worker(i, tasks, results, &wg)
// 	}
// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()
// 	// seding task
// 	for i := range numWorkers {
// 		tasks <- i
// 	}
// 	close(tasks)

// 	for result := range results {
// 		fmt.Println("Result :", result)
// 	}

// }

///=======simple WAITGROUP  WITHOUT USING CHANNEL========
// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	//wg.Add(1)  // wrong practice
// 	fmt.Printf("Working %d  starting \n", id)
// 	time.Sleep(time.Second)
// 	fmt.Printf("worker %d finished", id)
// }
// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3

// 	wg.Add(numWorkers) //THIS IS THE CORRECT WAY OF ADDING COUNTER TO WAITGROUP

// 	// Launch worker
// 	for i := range numWorkers {
// 		go worker(i, &wg)
// 	}
// 	wg.Wait()
// 	fmt.Println("All worker finished")
// }
