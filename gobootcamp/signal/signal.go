package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Get the current process ID.
	// Useful when sending signals from another terminal.
	pid := os.Getpid()
	fmt.Println("Process ID:", pid)

	// Create a buffered channel to receive OS signals.
	// Buffer size 1 prevents signal delivery from immediately
	// blocking if the program is temporarily busy.
	sigs := make(chan os.Signal, 1)

	// Register the signals this program wants to receive.
	//
	// SIGINT  : Interrupt (Ctrl+C).
	// SIGTERM : Termination request.
	// SIGHUP  : Terminal hangup / reload convention.
	// SIGUSR1 : User-defined signal 1.
	signal.Notify(
		sigs,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
		syscall.SIGUSR1,
	)

	// Stop signal notifications when main returns.
	defer signal.Stop(sigs)

	// Ticker generates an event every second.
	// Unlike time.Sleep inside a loop, a ticker works
	// naturally with select and other channel operations.
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	fmt.Println("Application started. Waiting for signals...")

	// Main event loop.
	for {
		select {
		// Receive OS signals.
		case sig := <-sigs:
			fmt.Println("\nReceived signal:", sig)

			switch sig {

			case syscall.SIGINT:
				fmt.Println("Received SIGINT (Interrupt)")
				fmt.Println("Gracefully shutting down...")

				return

			case syscall.SIGTERM:
				fmt.Println("Received SIGTERM (Terminate)")
				fmt.Println("Gracefully shutting down...")

				return

			case syscall.SIGHUP:
				fmt.Println("Received SIGHUP (Hangup)")
				fmt.Println("Configuration reload requested")

				// Place configuration reload logic here.

			case syscall.SIGUSR1:
				fmt.Println("Received SIGUSR1 (User-defined signal)")
				fmt.Println("Executing custom user-defined function")

				// Place custom functionality here.
			}

		// Perform periodic background work.
		case <-ticker.C:
			fmt.Println("Working...")
		}
	}

	// Returning from main ends the application.
	// Deferred cleanup functions are executed automatically.
}
