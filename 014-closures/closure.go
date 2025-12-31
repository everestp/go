package main

import "fmt"

// counter function returns another function (closure)
func counter() func() int {
	var count int = 0 // this variable is "captured" by the closure
	return func() int {
		count++        // increment the captured variable
		return count   // return the updated value
	}
}

func main() {
	// Create a closure instance
	increment := counter() // increment now holds the inner function with its own 'count'

	fmt.Println(increment()) // calls the closure, count becomes 1, prints 1
	fmt.Println(increment()) // calls the closure again, count becomes 2, prints 2

	// If you create another closure, it will have its own separate count
	newCounter := counter()
	fmt.Println(newCounter()) // prints 1
	fmt.Println(newCounter()) // prints 2
}
