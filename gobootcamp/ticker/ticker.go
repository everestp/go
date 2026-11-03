package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	stop := time.After(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case tick := <-ticker.C:
			fmt.Println("Tick at :", tick)
		case <-stop:
			fmt.Println("Stoping ticker")
			return

		}
	}
}

// func periodicTask() {
// 	fmt.Println("Performing periodic task", time.Now())
// }

// func main() {
// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()
// 	for {
// 		select {
// 		case <-ticker.C:
// 			periodicTask()
// 		}
// 	}
// }

// func main() {
// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()
// 	for tick := range 5 {
// 		fmt.Println("Tick at:", tick)
// 	}
// 	for tick := range ticker.C {
// 		fmt.Println("Tick at:", tick)
// 	}

// }
