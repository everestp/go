package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// User struct uses 'Tags' to map Go fields to JSON keys.
// Note: Fields MUST be capitalized (Exported) to be visible to the json package.
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// 1. DATA PREPARATION
	user := User{Name: "Everest", Email: "everest@cka.one"}
	fmt.Println("Original Struct:", user)

	// ==========================================================
	// PART A: MARSHAL & UNMARSHAL (Memory Based)
	// Use these when you already have the full data in a variable.
	// ==========================================================

	// Marshal: Convert Struct -> []byte (JSON)
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal("Marshal Error:", err)
	}
	fmt.Println("1. Marshal (Struct to String):", string(jsonData))

	// Unmarshal: Convert []byte (JSON) -> Struct
	// IMPORTANT: You must pass a pointer (&user1) so the function can fill it.
	var user1 User
	err = json.Unmarshal(jsonData, &user1)
	if err != nil {
		log.Fatal("Unmarshal Error:", err)
	}
	fmt.Println("2. Unmarshal (String to Struct):", user1)

	// ==========================================================
	// PART B: ENCODER & DECODER (Stream Based)
	// Use these when reading/writing directly to a Connection, File, or Buffer.
	// Very efficient for large data or HTTP Request/Response bodies.
	// ==========================================================

	// NewDecoder: Reading from a Stream (strings.Reader simulates a file or connection)
	jsonData1 := `{"name":"Cka","email":"Cka@cka.one"}`
	reader := strings.NewReader(jsonData1) // Implements io.Reader
	decoder := json.NewDecoder(reader)

	var user2 User
	// Decode() reads the stream and fills the struct directly
	err = decoder.Decode(&user2)
	if err != nil {
		log.Fatal("Decode Error:", err)
	}
	fmt.Println("3. Decoder (From Stream):", user2)

	// NewEncoder: Writing to a Stream (bytes.Buffer simulates a file or connection)
	var buf bytes.Buffer // Implements io.Writer
	encoder := json.NewEncoder(&buf)
	
	// Encode() takes the struct and writes JSON directly to the buffer
	err = encoder.Encode(user)
	if err != nil {
		log.Fatal("Encode Error:", err)
	}
	fmt.Printf("4. Encoder (To Stream): %s", buf.String())
}