package main

import (
	"fmt"
	"time"
)

func channel() {
	ch := make(chan int)

	for i := range 5 {
		go func(i int) {
			time.Sleep(time.Second)
			fmt.Println("Worker finished:", i)
			ch <- i
		}(i)
	}

	for range 5 {
		value := <-ch
		fmt.Println("Main received:", value)
	}

	fmt.Println("All workers sent their values. Main finished.")
}
