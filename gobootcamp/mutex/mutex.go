package main

import (
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	numGoroutines := 5
	wg.Add(numGoroutines)
	increamen := func() {
		for range 1000 {
			mu.Lock()
			counter++
			mu.Unlock()

		}

	}

	for range numGoroutines {
		go increamen()
	}

	wg.Wait()

}

// type counter struct {
// 	mu   sync.Mutex
// 	cout int
// }

// func (c *counter) increament() {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.cout++

// }

// func (c *counter) getValue() int {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.cout
// }

// func main() {
// 	var wg sync.WaitGroup
// 	counter := &counter{}

// 	numGoroutines := 10

// 	for range numGoroutines {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for range 100 {
// 				counter.increament()

// 			}
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println("Final value :", counter.getValue())

// }
