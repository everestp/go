package main

import (
	"fmt"
	"runtime"
	"sync"
)

func HeavyTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task %d is starting \n", id)
	for range 100_000_000 {

	}
	fmt.Println("task finsidhed", id)

}

func main() {
	numThreads := 4
	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup
	for i := range numThreads {
		wg.Add(1)
		go HeavyTask(i, &wg)
	}
	wg.Wait()
}
