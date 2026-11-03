package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		log.Fatalln("Error reading form reader")
	}
	fmt.Println("", string(buf[:n]))
}

func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error writing from reader ")
	}
}

func closeResource(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatalln("Error reading form reader")
	}
}

func bufferExample() {
	var buf bytes.Buffer //create memory in stack
	buf.WriteString("Hello buffer")
	fmt.Println(buf.String())
}

func multiReaderExample() {
	r1 := strings.NewReader("Hello  ")
	r2 := strings.NewReader("Worlds  ")
	mr := io.MultiReader(r1, r2)
	buf := new(bytes.Buffer) //create mempory on heap
	_, err := buf.ReadFrom(mr)
	if err != nil {
		log.Fatalln("Error reading form reader")
	}
	fmt.Println(buf.String())
}

func writeTofile(filepath string, data string) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Error opening creaing fike")
	}
	defer closeResource(file)

	// _, err = file.Write([]byte(data))
	// if err != nil {
	// 	log.Fatalln("Error reading form reader")
	// }
	//Alternate way of wrting a file
	//Type(value)
	writer := io.Writer(file)
	_, err = writer.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error opening creaing fike")
	}
}

func pipeExample() {
	pr, pw := io.Pipe()
	go func() {
		pw.Write([]byte("Hello pipe"))
		pw.Close()
	}()
	buf := new(bytes.Buffer)
	buf.ReadFrom(pr)
	fmt.Println(buf.String())

}

func main() {
	fmt.Println("===Read from reader===")
	readFromReader(strings.NewReader("Hello Reader!"))
	var writer bytes.Buffer
	fmt.Println("===write from writer===")
	writeToWriter(&writer, "Hello writer")
	fmt.Println(writer.String())

	fmt.Println("===Multi Reader Example===")
	multiReaderExample()
	fmt.Println("===Pipe Example===")
	pipeExample()
	filepath := "io.txt"
	writeTofile(filepath, "Hello this is the file write")
}
