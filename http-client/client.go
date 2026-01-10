package main

import (
	"encoding/json" // For parsing JSON
	"fmt"
	"net/http"      // For the HTTP client
	"time"          // Good practice to add timeouts
)

// Post represents the structure of the JSON returned by the API
// Field names MUST be capitalized to be "Exported" (visible to the json package)
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	// 1. BEST PRACTICE: Always define a timeout for your client.
	// A default http.Client{} has no timeout and can hang forever.
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// 2. Make the GET request
	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Printf("❌ Error making GET Request: %v\n", err)
		return
	}

	// 3. IMPORTANT: Always close the body to prevent memory leaks.
	// defer ensures it runs even if the function exits early.
	defer resp.Body.Close()

	// 4. Check the Status Code
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("⚠️ Server returned status: %d\n", resp.StatusCode)
		return
	}

	// 5. DECODING (Streaming approach)
	// Instead of reading all bytes into memory first (io.ReadAll), 
	// we decode the stream directly into our struct.
	var myPost Post
	err = json.NewDecoder(resp.Body).Decode(&myPost)
	if err != nil {
		fmt.Printf("❌ Error decoding JSON: %v\n", err)
		return
	}

	// 6. Access your data
	fmt.Println("✅ Successfully fetched post:")
	fmt.Printf("Title: %s\n", myPost.Title)
	fmt.Printf("Body:  %s\n", myPost.Body)
}