package main

import (
	"fmt"
	"time"
)

// Helper function for the periodic task example
func periodicTask() {
	fmt.Println("-> Performing periodic task at:", time.Now().Format("15:04:05"))
}

func main() {
	// ==========================================
	// 1. TICKER WITH TIMEOUT (time.After + select)
	// ==========================================
	fmt.Println("=== 1. TICKER WITH TIMEOUT ===")
	ticker1 := time.NewTicker(time.Second)
	stop := time.After(3 * time.Second) // Set timeout for 3 seconds

	// Always defer Stop() to prevent ticker memory leaks
	defer ticker1.Stop()

	for {
		select {
		case tick := <-ticker1.C:
			fmt.Println("Tick at:", tick.Format("15:04:05"))
		case <-stop:
			fmt.Println("Stopping ticker after timeout...\n")
			goto NEXT_PATTERN // Break out of this loop to move to the next pattern
		}
	}

NEXT_PATTERN:

	// ==========================================
	// 2. PERIODIC TASK RUNNER (Infinite loop)
	// ==========================================
	fmt.Println("=== 2. PERIODIC TASK RUNNER (Demo limit: 2 ticks) ===")
	ticker2 := time.NewTicker(time.Second)
	defer ticker2.Stop()

	count := 0
	for {
		select {
		case <-ticker2.C:
			periodicTask()
			count++
			// Breaking early so the script can continue to the next example
			if count >= 2 {
				fmt.Println("Periodic demo finished.\n")
				goto NEXT_PATTERN_2
			}
		}
	}

NEXT_PATTERN_2:

	// ==========================================
	// 3. RANGE LOOPS (Range-over-int & Ticker Range)
	// ==========================================
	fmt.Println("=== 3. RANGE-OVER-INT EXAMPLE ===")
	// Go 1.22+ feature: range over an integer (loops 5 times, i goes from 0 to 4)
	for i := range 5 {
		fmt.Printf("Range tick index: %d\n", i)
		time.Sleep(200 * time.Millisecond)
	}

	fmt.Println("\n=== ALL TICKER PATTERNS FINISHED ===")
}
