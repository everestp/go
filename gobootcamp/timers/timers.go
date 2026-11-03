package main

import (
	"fmt"
	"time"
)

//=====SCHEDULING DELAYED OPERATIO

func main() {
	timer := time.NewTimer(2 * time.Second) // non blocking timer start
	go func() {
		<-timer.C
		fmt.Println("Delayed operation executd")
	}()
	fmt.Println("Waiting")
	time.Sleep(3 * time.Second) // blocking timers starts
	fmt.Println("End")

}

//====== TIMEOUT =====
// func longRunning() {
// 	for i := range 20 {
// 		fmt.Println(i)
// 		time.Sleep(time.Second)

// 	}
// }

// func main() {
// 	timeout := time.After(5 * time.Second)
// 	done := make(chan bool)
// 	go func() {
// 		longRunning()
// 		done <- true
// 	}()
// 	select {
// 	case <-timeout:
// 		fmt.Println("Operation times out")
// 	case <-done:
// 		fmt.Println("Operatiion time out ")

// 	}
// }

//========BASIC TIMER  USER =====
// func main() {

// 	timer := time.NewTimer(2 * time.Second)
// 	fmt.Println("Waiting for timer.c")
// 	stopped := timer.Stop()
// 	if stopped {
// 		fmt.Println("Timer stopped")
// 	}
// 	timer.Reset(time.Second)

// 	<-timer.C //blocking nature
// 	timer.Stop()
// 	fmt.Println("Timer expired")
// }
