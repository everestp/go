// Package main demonstrates the use of interfaces, type assertions, and custom error handling in Go
package main

import (
	"fmt"
	"math"
)

// --------------------------
// Interface Definitions
// --------------------------

// Shape interface defines a contract for any geometric shape that can calculate its area.
// Any type implementing Area() float64 satisfies this interface.
type Shape interface {
	Area() float64
}

// Measurable interface defines a contract for shapes that can calculate perimeter.
type Measurable interface {
	Perimetre() float64
}

// Geometry interface embeds both Shape and Measurable.
// Any type implementing both Area() and Perimetre() methods satisfies this interface.
type Geometry interface {
	Shape
	Measurable
}

// --------------------------
// Struct Definitions
// --------------------------

// Rectangle struct represents a rectangle with length and breadth.
type Rectangle struct {
	breadth float64
	length  float64
}

// Circle struct represents a circle with a radius.
type Circle struct {
	radius float64
}

// --------------------------
// Method Implementations
// --------------------------

// Perimetre calculates the perimeter of the rectangle.
// Satisfies Measurable interface.
func (r Rectangle) Perimetre() float64 {
	return 2 * (r.breadth + r.length)
}

// Area calculates the area of the rectangle.
// Satisfies Shape interface.
func (r Rectangle) Area() float64 {
	return r.breadth * r.length
}

// Area calculates the area of the circle.
// Satisfies Shape interface.
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// --------------------------
// Functions Using Interfaces
// --------------------------

// calculateArea accepts any type that satisfies the Shape interface and returns its area.
func calculateArea(shape Shape) float64 {
	return shape.Area()
}

// desscribeShape accepts any type that satisfies Geometry (both Shape & Measurable)
// and prints both area and perimeter.
func desscribeShape(g Geometry) {
	fmt.Println("Area       :=", g.Area())
	fmt.Println("Perimeter  :=", g.Perimetre())
}

// descibeValue accepts an empty interface, which means it can take **any type**.
// Demonstrates type assertion and type inspection.
func descibeValue(t interface{}) {
	fmt.Printf("Type : %T, Value : %v\n", t, t)
}

// --------------------------
// Custom Error Handling
// --------------------------

// CaculationError defines a custom error type with a message.
// Implements the error interface by providing an Error() string method.
type CaculationError struct {
	msg string
}

// Error method satisfies the built-in error interface.
func (ce CaculationError) Error() string {
	return ce.msg
}

// performCalculation performs a calculation (here, square root).
// Returns an error if input is invalid (negative value).
func performCalculation(val float64) (float64, error) {
	if val < 0 {
		// Return zero value and a custom error
		return 0, CaculationError{msg: "Invalid input: cannot calculate square root of negative number"}
	}
	// Return square root and nil error
	return math.Sqrt(val), nil
}

// --------------------------
// Main Function
// --------------------------
func main() {
	// Create rectangle instances
	rect := Rectangle{2, 3}
	rect1 := Rectangle{2, 10}

	// Calculate area using Shape interface
	fmt.Println("The area of rect is:", calculateArea(rect))

	// Describe rectangles using Geometry interface
	// Note: Only types that implement both Area() and Perimetre() can be passed
	desscribeShape(rect)
	desscribeShape(rect1)

	// --------------------------
	// Using empty interface and type assertion
	// --------------------------
	mystrtBox := interface{}(10) // empty interface can hold any type
	descibeValue(mystrtBox)     // prints type and value

	// Type assertion to retrieve the original int value
	retriveInt, ok := mystrtBox.(int)
	if ok {
		fmt.Println("Retrieved int:", retriveInt)
	} else {
		fmt.Println("Value is not an integer")
	}

	// --------------------------
	// Demonstrating custom error handling
	// --------------------------
	val := -9.0
	result, err := performCalculation(val)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Square root:", result)
	}
}
