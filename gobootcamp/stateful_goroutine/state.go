package main

import (
	"fmt"
	"time"
)

// StatefulWorker maintains its own state (count).
// Only the worker goroutine modifies count, avoiding concurrent writes.
type StatefulWorker struct {
	count int
	ch    chan int // Channel used to receive values from the sender.
}

// Start launches a goroutine that continuously listens for incoming values.
//
// The select statement waits for a channel operation to become ready.
// Whenever a value arrives, the worker adds it to count.
func (w *StatefulWorker) Start() {
	go func() {
		for {
			select {
			case value := <-w.ch:
				w.count += value
				fmt.Println("Current count:", w.count)
			}
		}
	}()
}

// Send sends a value to the worker through its channel.
//
// Since this is an unbuffered channel, Send blocks until a receiver
// is ready to receive the value.
func (w *StatefulWorker) Send(value int) {
	w.ch <- value
}

func main() {
	// Create a worker with an unbuffered channel.
	stWorker := &StatefulWorker{
		ch: make(chan int),
	}

	// Start the worker goroutine before sending values.
	stWorker.Start()

	// Send five values (0, 1, 2, 3, 4) to the worker.
	for i := range 5 {
		stWorker.Send(i)
		time.Sleep(5 * time.Millisecond)
	}
}
