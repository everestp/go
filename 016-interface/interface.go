package main

import "fmt"

// ----------------------
// Define the interface
// ----------------------

// paymenter is an interface that defines a single behavior: pay
// Any type that has a method `pay(amount float32)` automatically implements this interface
type paymenter interface {
	pay(amount float32)
	refund(amount float32)
}

// ----------------------
// Define the struct that uses the interface
// ----------------------

// payment struct represents a payment service
// It does NOT implement paymenter itself
// It uses a field `gateway` which can be any type that implements paymenter
type payment struct {
	gateway paymenter // this is the "skill" or "worker" that knows how to pay
}

// makePayment is a method on payment struct
// This method delegates the actual payment to the gateway
func (p payment) makePayment(amount float32) {
	// call the pay method of the injected gateway
	p.gateway.pay(amount)
}

// ----------------------
// Define a concrete payment gateway
// ----------------------

// razorpay struct represents the Razorpay payment gateway
type razorpay struct {
	// we can add fields like API key, secret, etc. in real scenario
}

// Implement the pay method for Razorpay
// Since razorpay has pay method, it implements paymenter interface
func (r razorpay) pay(amount float32) {
	// logic to make payment (simplified)
	fmt.Println("Making payment using Razorpay:", amount)
}
func (r razorpay) refund(amount float32) {
	// logic to make payment (simplified)
	fmt.Println("Making refund using Razorpay:", amount)
}
// ----------------------
// (Optional) Another payment gateway
// ----------------------
// You can add another gateway like Stripe easily:

type stripe struct {
	// Stripe-specific fields
}

func (s stripe) pay(amount float32) {
	//l
	fmt.Println("Making payment using Stripe:", amount)
}
func (s stripe) refund(amount float32) {
	//l
	fmt.Println("Making refund using Stripe:", amount)
}

// ----------------------
// Main function
// ----------------------
func main() {
	// Create a new payment using Razorpay as the gateway
	newPayment := payment{
		gateway: razorpay{}, // inject Razorpay as the "worker"
	}

	// Make a payment of 100
	// The payment struct delegates the actual payment to the gateway
	newPayment.makePayment(100)

	// If you want to switch to Stripe, just inject it here (no change to payment struct)
	newPayment.gateway = stripe{}
	newPayment.makePayment(200)
	newPayment.gateway.refund(30)
}
