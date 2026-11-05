package main

import (
	"fmt"
	"time"
)

type tickerRequest struct {
	personID   int
	numTickets int
	cost       int
}

//simulate processing fo ticker request

func ticketProcessing(requests <-chan tickerRequest, results chan<- int) {
	for req := range requests {
		fmt.Printf("Processing %d ticket(s) of personID %d with total cost %d \n", req.numTickets, req.personID, req.cost)
		//simulrate processing time
		time.Sleep(time.Second)
		results <- req.personID
	}
}
func main() {
	numofWorker := 1
	numofJobs := 10
	reqests := make(chan tickerRequest, numofJobs)
	results := make(chan int, numofJobs)

	// create a workes
	for range numofWorker {
		go ticketProcessing(reqests, results)
	}

	// send work to worker through channel

	for i := range numofJobs {
		reqests <- tickerRequest{
			personID:   i,
			numTickets: i * 2,
			cost:       9 + 3 + i,
		}
	}

	// get back the resilt form worker
	for range numofJobs {
		result := <-results
		fmt.Println("This is the result", result)
	}
}

// func worker(id int, tasks <-chan int, results chan<- int) {
// 	for task := range tasks {
// 		fmt.Printf("Worker  %d processing task %d \n", id, task)
// 		//Simulate work
// 		time.Sleep(time.Second)
// 		results <- task * 2
// 	}
// }

// func main() {
// 	numWorkers := 3
// 	numJobs := 10
// 	tasks := make(chan int, numJobs)
// 	results := make(chan int, numJobs)

// 	//create workers
// 	for i := range numWorkers {
// 		go worker(i, tasks, results)

// 	}
// 	// send values to the  task channel
// 	for i := range numJobs {
// 		tasks <- i
// 	}
// 	close(tasks)

// 	// collect the results
// 	for range numJobs {
// 		result := <-results
// 		fmt.Println("Result :", result)
// 	}

// }
