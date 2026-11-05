package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings" // uncomment if you use strings.NewReader for the reader example
)

func main() {
	// ============================================================================
	// 1. BUFIO READER EXAMPLE (Commented out for reference)
	// ============================================================================
	/*
		reader := bufio.NewReader(strings.NewReader("Hello bufio package\n i am here"))

		// Reading into a byte slice buffer
		data := make([]byte, 20)
		n, err := reader.Read(data)
		if err != nil {
			fmt.Println("Error reading:", err)
		}
		fmt.Printf("Read %d bytes: %s\n", n, data[:n])

		// Reading until a delimiter (e.g., newline '\n')
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading line:", err)
		}
		fmt.Print("Read line: ", line)
	*/

	// ============================================================================
	// 2. BUFIO WRITER EXAMPLE
	// ============================================================================
	// bufio.Writer adds an internal buffer layer over an io.Writer (like os.Stdout).
	// This reduces system calls by holding data in memory until it's flushed.
	writer := bufio.NewWriter(os.Stdout)

	// Writing a byte slice to the buffer
	data := []byte("Hello bufio package\n")
	n, err := writer.Write(data)
	if err != nil {
		fmt.Println("Error writing bytes:", err)
		return
	}
	fmt.Printf("Buffered %d bytes\n", n)

	// CRITICAL: Flush pushes the internal memory buffer out to the actual underlying destination (os.Stdout)
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing buffer:", err)
		return
	}

	// Writing a standard string directly using bufio helper methods
	str := "This is a bufio string\n"
	n, err = writer.WriteString(str)
	if err != nil {
		fmt.Println("Error writing string:", err)
		return
	}
	fmt.Printf("Buffered string of length %d\n", n)

	// Flush again to make sure the final string prints out to the screen
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing buffer:", err)
		return
	}
}
