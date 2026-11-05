package main

import (
	"fmt"
	"time"
)

func buffer_channel() {
	// =====BLOCKING ON RECEIVE ONLY IF BUFFER IS EMPTY
	ch := make(chan int, 2)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
		ch <- 2
	}()
	fmt.Println("Value :", <-ch)
	fmt.Println("Value :", <-ch)
	fmt.Println("End of program :")

}

// func main() {
//==============================Blocking ON SEND ONLY IF BUFFER  IS FULL
// 	// make(chan type, capacity)
// 	ch := make(chan int, 2)
// 	ch <- 1
// 	ch <- 2
// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		fmt.Println("Recieved :", <-ch)
// 	}()
// 	fmt.Println("Blocking starts :")
// 	ch <- 3 // Blocking because buffer is full
// 	fmt.Println("Blocking ends :")
// 	fmt.Println("Recieved :", <-ch)
// 	fmt.Println("Recieved :", <-ch)

// 	fmt.Println("Buffered channel")

// }
