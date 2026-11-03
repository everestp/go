package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// ============================================================================
// 1. CORE INTERFACE HELPERS (io.Reader, io.Writer, io.Closer)
// ============================================================================

// readFromReader demonstrates how to read data from any io.Reader.
// It requires an "empty bucket" (byte slice) to catch the incoming stream.
func readFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		log.Fatalln("Error reading from reader:", err)
	}
	fmt.Println("Read output:", string(buf[:n]))
}

// writeToWriter demonstrates how to write data into any io.Writer.
// (e.g., files, buffers, network connections, or pipes).
func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error writing to writer:", err)
	}
}

// closeResource demonstrates how to close any resource that implements io.Closer.
func closeResource(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatalln("Error closing resource:", err)
	}
}

// ============================================================================
// 2. BUFFER & MULTI-READER EXAMPLES
// ============================================================================

func bufferExample() {
	var buf bytes.Buffer // In-memory buffer storage
	buf.WriteString("Hello buffer")
	fmt.Println(buf.String())
}

// multiReaderExample chains multiple readers sequentially into a single stream.
func multiReaderExample() {
	r1 := strings.NewReader("Hello ")
	r2 := strings.NewReader("World! ")

	// Combine r1 and r2 so they read back-to-back
	mr := io.MultiReader(r1, r2)

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(mr)
	if err != nil {
		log.Fatalln("Error reading from multi-reader:", err)
	}
	fmt.Println("MultiReader result:", buf.String())
}

// ============================================================================
// 3. PIPES & FILE I/O EXAMPLES
// ============================================================================

// pipeExample demonstrates an synchronous in-memory pipe.
// The writer runs in a goroutine because io.Pipe is unbuffered and blocks
// until a reader is actively listening.
func pipeExample() {
	pr, pw := io.Pipe()

	// Writer goroutine
	go func() {
		pw.Write([]byte("Hello synchronized pipe!"))
		pw.Close() // Signal EOF to the reader
	}()

	// Reader (main thread) pulls everything into a buffer
	buf := new(bytes.Buffer)
	buf.ReadFrom(pr)

	fmt.Println("Pipe output:", buf.String())
}

// writeTofile opens or creates a file and writes text data to it using an io.Writer interface.
func writeTofile(filepath string, data string) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Error opening/creating file:", err)
	}

	// Ensures the file is safely closed when the function exits
	defer closeResource(file)

	// Treat the file as a generic io.Writer interface
	var writer io.Writer = file
	_, err = writer.Write([]byte(data + "\n"))
	if err != nil {
		log.Fatalln("Error writing to file:", err)
	}
	fmt.Println("Successfully wrote data to:", filepath)
}

// ============================================================================
// 4. CUSTOM TYPE IMPLEMENTING io.Closer
// ============================================================================

type MyResource struct {
	name string
}

// Close satisfies the io.Closer interface for MyResource
func (m MyResource) Close() error {
	fmt.Println("Closing custom resource:", m.name)
	return nil
}

// ============================================================================
// 5. MAIN ENTRY POINT
// ============================================================================

func main() {
	fmt.Println("=== 1. Read / Write Interfaces ===")
	readFromReader(strings.NewReader("Hello Reader!"))

	var writer bytes.Buffer
	writeToWriter(&writer, "Hello writer")
	fmt.Println("Writer output:", writer.String())

	fmt.Println("\n=== 2. Multi Reader Example ===")
	multiReaderExample()

	fmt.Println("\n=== 3. Pipe Example ===")
	pipeExample()

	fmt.Println("\n=== 4. File I/O Example ===")
	writeTofile("io.txt", "Hello, this is a file write test!")

	fmt.Println("\n=== 5. Custom Closer Example ===")
	resource := &MyResource{
		name: "Mount Everest Connection",
	}
	closeResource(resource)
}
