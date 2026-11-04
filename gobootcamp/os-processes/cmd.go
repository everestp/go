package main

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"
)

func main() {
	// 1. Execute a command and capture stdout + stderr.
	runCombinedOutput()

	// 2. Execute a command and capture stdout only.
	runOutput()

	// 3. Pass input to a command using strings.Reader.
	runWithStringReader()

	// 4. Pass input to a command using io.Pipe.
	runWithPipe()

	// 5. Read an environment variable using a command.
	readEnvironmentVariable()

	// 6. Start a long-running process, kill it, and wait.
	startAndKillProcess()
}

// CombinedOutput executes a command and captures both stdout and stderr.
//
// Use this when you want all command output in one byte slice,
// including error messages written to stderr.
func runCombinedOutput() {
	fmt.Println("\n--- CombinedOutput ---")

	cmd := exec.Command("ls", "-l")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(output))
}

// Output executes a command and captures only stdout.
//
// If the command writes to stderr, Output does not include that data.
// Output automatically starts the command and waits for it to finish.
func runOutput() {
	fmt.Println("\n--- Output ---")

	cmd := exec.Command("echo", "Hello World!")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(output))
}

// Stdin allows us to provide input to a process.
//
// strings.NewReader implements io.Reader, so the command can
// consume the provided string as its standard input.
//
// Here grep searches for lines containing "foo".
func runWithStringReader() {
	fmt.Println("\n--- Stdin with strings.Reader ---")

	cmd := exec.Command("grep", "foo")

	cmd.Stdin = strings.NewReader(
		"food is good\nbar\nbaz\n",
	)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(output))
}

// io.Pipe connects a writer and reader.
//
// Anything written to pw can be read from pr.
// We assign pr to the command's Stdin.
//
// The goroutine writes data into the pipe while grep
// reads that data and searches for "foo".
func runWithPipe() {
	fmt.Println("\n--- Stdin with io.Pipe ---")

	pr, pw := io.Pipe()

	cmd := exec.Command("grep", "foo")
	cmd.Stdin = pr

	// Start a goroutine to write data into the pipe.
	go func() {
		defer pw.Close()

		_, err := pw.Write(
			[]byte("food is good\nbar\nbaz\n"),
		)

		if err != nil {
			fmt.Println("Pipe write error:", err)
		}
	}()

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(output))
}

// Execute printenv to retrieve an environment variable.
//
// printenv SHELL prints the value of the SHELL environment variable.
func readEnvironmentVariable() {
	fmt.Println("\n--- Environment Variable ---")

	cmd := exec.Command("printenv", "SHELL")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("SHELL:", strings.TrimSpace(string(output)))
}

// Start launches a process without waiting for it to finish.
//
// Process.Kill terminates the running process.
//
// Wait must be called after Start to release resources
// associated with the process and collect its exit status.
func startAndKillProcess() {
	fmt.Println("\n--- Start, Kill and Wait ---")

	cmd := exec.Command("sleep", "60")

	// Start the process asynchronously.
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting process:", err)
		return
	}

	fmt.Println("Process started. PID:", cmd.Process.Pid)

	// Allow the process to run for two seconds.
	time.Sleep(2 * time.Second)

	// Terminate the running process.
	err = cmd.Process.Kill()
	if err != nil {
		fmt.Println("Error killing process:", err)
		_ = cmd.Wait()
		return
	}

	fmt.Println("Process killed")

	// Wait for the process to exit and release its resources.
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Process exited with:", err)
		return
	}

	fmt.Println("Process completed successfully")
}
