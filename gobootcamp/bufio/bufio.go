package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// reader := bufio.NewReader(strings.NewReader("Hello bufio packagef bb\n i am here"))

	// // Reading bytes slice
	// data := make([]byte, 20)
	// n, err := reader.Read(data)
	// if err != nil {
	// 	fmt.Println("Error reading", err)

	// }
	// fmt.Printf("Read %d bytes : %s", n, data[:n])
	// line, err := reader.ReadString('\n')
	// fmt.Println(line)

	writer := bufio.NewWriter(os.Stdout)
	// writing byte slicer
	data := []byte("Hello bufio package")
	n, err := writer.Write(data)
	if err != nil {
		fmt.Println("Error writing", err)
		return
	}
	fmt.Println("Wrote byte", n)
	//Flush thge buffer to ensue all data is written to os.Stdout
	err = writer.Flush()
	if err != nil {
		fmt.Println("Errir flushing")
		return
	}

	str := "This is the string \n"
	n, err = writer.WriteString(str)
	if err != nil {
		fmt.Println("Error writing", err)
	}
	fmt.Println("Wrote byte", n)
	err = writer.Flush()
	if err != nil {
		fmt.Println("Errir flushing")
	}
}
