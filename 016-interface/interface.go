package main

import "fmt"

// ----------------------
// Define the interface
// ----------------------

// paymenter is an interface that defines a single behavior: pay and refund
// Any type that has these methods automatically implements this interface
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

// ----------------------
// Dependency Inversion happens here!
// ----------------------
// The payment struct (high-level module) depends on the abstraction `paymenter` (interface)
// rather than a concrete implementation like Razorpay or Stripe.
// This allows us to inject any gateway that implements paymenter.
// High-level module: payment
// Abstraction: paymenter
// Low-level modules: razorpay, stripe

// makePayment is a method on payment struct
// This method delegates the actual payment to the gateway
func (p payment) makePayment(amount float32) {
	fmt.Println("Payment service: initiating payment of", amount)
	// call the pay method of the injected gateway
	p.gateway.pay(amount)
}

// makeRefund is another method to handle refunds
func (p payment) makeRefund(amount float32) {
	fmt.Println("Payment service: initiating refund of", amount)
	// call the refund method of the injected gateway
	p.gateway.refund(amount)
}

// ----------------------
// Define a concrete payment gateway
// ----------------------

// razorpay struct represents the Razorpay payment gateway
type razorpay struct {
	// we can add fields like API key, secret, etc. in real scenario
}

// Implement the pay method for Razorpay
func (r razorpay) pay(amount float32) {
	fmt.Println("Razorpay gateway: processing payment of", amount)
}

// Implement the refund method for Razorpay
func (r razorpay) refund(amount float32) {
	fmt.Println("Razorpay gateway: processing refund of", amount)
}

// ----------------------
// Another concrete payment gateway (Stripe)
// ----------------------

type stripe struct {
	// Stripe-specific fields
}

func (s stripe) pay(amount float32) {
	fmt.Println("Stripe gateway: processing payment of", amount)
}

func (s stripe) refund(amount float32) {
	fmt.Println("Stripe gateway: processing refund of", amount)
}

// ----------------------
// Main function
// ----------------------
func main() {
	// ----------------------
	// Example 1: Use Razorpay
	// ----------------------

	// Create a new payment using Razorpay as the gateway
	newPayment := payment{
		gateway: razorpay{}, // injecting the dependency (Razorpay implements paymenter)
	}

	// Make a payment of 100
	// Payment struct delegates the actual payment to the gateway
	newPayment.makePayment(100)

	// Make a refund of 20
	newPayment.makeRefund(20)

	// ----------------------
	// Example 2: Switch to Stripe
	// ----------------------

	// Without changing the payment struct or makePayment logic, we can swap the gateway
	newPayment.gateway = stripe{} // inject Stripe instead of Razorpay

	// Make a payment using Stripe
	newPayment.makePayment(200)

	// Make a refund using Stripe
	newPayment.makeRefund(50)

	// ----------------------
	// ✅ Key concepts in this code
	// ----------------------

	// 1. Dependency Inversion Principle (DIP):
	//    - High-level module (payment) depends on abstraction (paymenter), not concrete implementations (Razorpay/Stripe).
	//    - Low-level modules (razorpay, stripe) implement the abstraction.
	//    - This allows us to switch gateways easily without changing payment struct code.

	// 2. Interfaces provide flexibility:
	//    - Any struct implementing paymenter can be injected.
	//    - Payment struct does not care about the concrete type.

	// 3. Composition over inheritance:
	//    - Payment struct "has a" gateway (composition), rather than "is a" gateway.
	//    - This avoids tight coupling and allows swapping behavior at runtime.

	// 4. Delegation:
	//    - Payment struct delegates payment and refund tasks to the injected gateway.
	//    - Makes it easy to add extra logic (logging, retry, validation) in payment struct.

	// 5. Extendable design:
	//    - To add a new gateway (e.g., PayPal), just implement paymenter.
	//    - No changes needed in payment struct.
}
