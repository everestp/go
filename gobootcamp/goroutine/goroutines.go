package main

import (
	"fmt"
	"time"
)

// Goroutine are just finction that leave the main thread and run in the back to join the main thread once the functios are finisjed/ready to return any value
// Goroutie do nit stop the program flow and are non-blocking
func main() {
	var err error
	fmt.Println("Beginning program")
	go sayHello()
	fmt.Println("After sayHello function")

	// err = go doWork()  // this is not accpeted
	go func() {

		err = doWork()
	}()

	go printNumbers()
	go printLetter()
	time.Sleep(4 * time.Second)

	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Print("Work done")
	}

}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("hello frm goroutine")

}
func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetter() {
	for _, letter := range "abcde" {
		fmt.Println(string(letter))
		time.Sleep(100 * time.Millisecond)
	}
}

func doWork() error {
	//Simualte work
	time.Sleep(1 * time.Second)
	return fmt.Errorf("Error occuered while doing work")
}
