package main

import (
	"fmt"
	"time"
)

func chanell() {
	ch := make(chan int)

	for i := range 50000 {
		go func() {
			time.Sleep(1 * time.Second)
			if i == 4999 {
				fmt.Println("I am not sending i ", i)
			}
			fmt.Println("Worker finished", i)
			// No send to ch!
		}()
	}

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Worker finished")
		// No send to ch!
	}()

	go func() {
		fmt.Println("Worker finished 2")
		// No send to ch!
	}()

	value := <-ch
	fmt.Println(value)
}
