package main

import (
	"fmt"
	"time"
)

// func main() {
// 	done := make(chan struct{})

// 	go func() {
// 		fmt.Print("Working")
// 		time.Sleep(2 * time.Second)
// 		done <- struct{}{}
// 	}()
// 	<-done
// 	fmt.Println("Finished")
// }

// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		ch <- 9 //Blocking until the value is received
// 		time.Sleep(2 * time.Second)
// 		fmt.Println("Sent value")

// 	}()
// 	value := <-ch
// 	fmt.Println(value)
// }

func main() {
	numGoroutines := 3
	done := make555555(chan int, 3)
	for i := range numGoroutines {
		go func(id int) {
			fmt.Printf("Goroutine %d working ...\n", id)
			time.Sleep(time.Second)
			done <- id

		}(i)
	}
	for range numGoroutines {
		<-done //Wait for each goroutine is finished

	}
	fmt.Println("All goroutine are finshed")
}
