package main

import (
	"fmt"
	"sync"
	"time"
)

// LeakyBucket controls how frequently requests are allowed.
//
// capacity  : Maximum number of available tokens.
// leakRate  : Time required to replenish one token.
// tokens    : Current available tokens.
// lastleak  : Last time at which token replenishment was calculated.
// mu        : Protects shared state from concurrent access.
type LeakyBucket struct {
	capacity int
	leakRate time.Duration
	tokens   int
	lastleak time.Time
	mu       sync.Mutex
}

// NewLeakyBucket initializes a new rate limiter.
//
// Initially, the bucket is full, meaning it has maximum available tokens.
// lastleak records the initialization time for future refill calculations.
func NewLeakyBucket(capacity int, leakRate time.Duration) *LeakyBucket {
	if capacity <= 0 {
		panic("capacity must be greater than zero")
	}

	if leakRate <= 0 {
		panic("leakRate must be greater than zero")
	}

	return &LeakyBucket{
		capacity: capacity,
		leakRate: leakRate,
		tokens:   capacity,
		lastleak: time.Now(),
	}
}

// Allow checks whether a request can be accepted.
//
// 1. Lock the mutex to protect shared state.
// 2. Calculate elapsed time since the last refill.
// 3. Determine how many tokens should be replenished.
// 4. Ensure tokens never exceed bucket capacity.
// 5. Update the last refill timestamp.
// 6. Consume one token if available.
//
// Returns true  : Request accepted.
// Returns false : Request rejected.
func (lb *LeakyBucket) Allow() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	// Get the current time.
	now := time.Now()

	// Calculate how much time has passed since the last refill.
	elapsedTime := now.Sub(lb.lastleak)

	// Calculate how many complete refill intervals have passed.
	//
	// Example:
	// elapsedTime = 1200ms
	// leakRate    = 500ms
	// tokensToAdd = 2
	tokensToAdd := int(elapsedTime / lb.leakRate)

	// Replenish the available tokens.
	lb.tokens += tokensToAdd

	// Prevent the bucket from exceeding its maximum capacity.
	if lb.tokens > lb.capacity {
		lb.tokens = lb.capacity
	}

	// Advance lastleak only by complete refill intervals.
	//
	// This preserves the remaining fractional elapsed time.
	// Updating it directly to now would discard that remainder.
	lb.lastleak = lb.lastleak.Add(
		time.Duration(tokensToAdd) * lb.leakRate,
	)

	// Display refill and consumption information.
	fmt.Printf(
		"Tokens added: %d | Tokens available: %d\n",
		tokensToAdd,
		lb.tokens,
	)

	fmt.Printf("Last refill time: %v\n", lb.lastleak)

	// Consume one token for the current request.
	if lb.tokens > 0 {
		lb.tokens--
		return true
	}

	// No tokens available: reject the request.
	return false
}

func main() {
	// Bucket capacity: 5 requests.
	// Refill interval: 500 milliseconds per token.
	leakyBucketInst := NewLeakyBucket(5, 500*time.Millisecond)

	// Simulate 10 incoming requests.
	for range 10 {
		if leakyBucketInst.Allow() {
			fmt.Println("Request accepted")
		} else {
			fmt.Println("Request rejected")
		}

		// Wait one second before sending the next request.
		time.Sleep(time.Second)
	}
}
