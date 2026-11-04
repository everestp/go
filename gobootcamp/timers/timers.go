package main

import (
	"fmt"
	"time"
)

func main() {
	// ==========================================
	// 1. BASIC TIMER CONTROL (Stop & Reset)
	// ==========================================
	fmt.Println("=== 1. BASIC TIMER CONTROL (Stop & Reset) ===")

	// Create a timer that is set to fire after 2 seconds
	timer := time.NewTimer(2 * time.Second)

	// Stop the timer immediately before it expires.
	// timer.Stop() returns true if the timer was successfully stopped.
	if timer.Stop() {
		fmt.Println("-> Timer was active and successfully stopped.")
	}

	// Reset the timer to fire after a shorter duration (1 second)
	timer.Reset(1 * time.Second)
	fmt.Println("-> Timer reset to 1 second. Waiting for expiration...")

	<-timer.C // Blocks until the timer channel receives the expiration signal
	fmt.Println("-> Timer expired!\n")

	// ==========================================
	// 2. SCHEDULING A DELAYED OPERATION
	// ==========================================
	fmt.Println("=== 2. SCHEDULING A DELAYED OPERATION ===")

	// Start a non-blocking timer for background execution
	delayedTimer := time.NewTimer(1 * time.Second)
	go func() {
		<-delayedTimer.C
		fmt.Println("-> [Async Background] Delayed operation executed successfully!")
	}()

	fmt.Println("-> Main thread continues instantly while background task waits...")
	time.Sleep(1200 * time.Millisecond) // Pause briefly to let the background goroutine finish printing
	fmt.Println()

	// ==========================================
	// 3. TIMEOUT PATTERN (select + time.After)
	// ==========================================
	fmt.Println("=== 3. TIMEOUT PATTERN (select + time.After) ===")

	// Create a timeout channel that fires after 2 seconds
	timeout := time.After(2 * time.Second)
	done := make(chan bool)

	// Simulate a long-running background task
	go func() {
		// This simulation takes 3 seconds (longer than our 2-second timeout)
		for i := 1; i <= 3; i++ {
			fmt.Printf("-> Working... step %d/3\n", i)
			time.Sleep(1 * time.Second)
		}
		done <- true // Signal completion
	}()

	// Use a select statement to race between the task finishing and the timeout expiring
	select {
	case <-timeout:
		fmt.Println("-> Result: Operation timed out! (Task took too long)")
	case <-done:
		fmt.Println("-> Result: Operation completed successfully within the time limit!")
	}

	fmt.Println("\n=== ALL TIMER PATTERNS FINISHED ===")
}
