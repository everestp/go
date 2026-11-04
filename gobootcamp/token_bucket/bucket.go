package main

import (
	"fmt"
	"time"
)

// RateLimiter implements a thread-safe Token Bucket rate limiter
// using Go channels and background tickers.
type RateLimiter struct {
	tokens     chan struct{} // Buffered channel acting as the token bucket
	refillRate time.Duration // Interval at which new tokens are added
}

// NewRateLimiter initializes and returns a new RateLimiter.
// - rateLimit: Maximum number of tokens the bucket can hold (burst capacity).
// - refillRate: How often a single token is replenished.
func NewRateLimiter(rateLimit int, refillRate time.Duration) *RateLimiter {
	rl := &RateLimiter{
		tokens:     make(chan struct{}, rateLimit),
		refillRate: refillRate,
	}

	// Pre-fill the bucket to its maximum capacity initially
	for i := 0; i < rateLimit; i++ {
		rl.tokens <- struct{}{}
	}

	// Start a background goroutine to periodically replenish tokens
	go rl.startRefill()

	return rl
}

// startRefill runs in the background, ticking at the specified refillRate
// to add a token back into the bucket without exceeding its capacity.
func (rl *RateLimiter) startRefill() {
	ticker := time.NewTicker(rl.refillRate)
	defer ticker.Stop()

	for range ticker.C {
		select {
		case rl.tokens <- struct{}{}:
			// Token successfully added
		default:
			// Bucket is already at maximum capacity; discard the extra token
		}
	}
}

// Allow checks if a request can proceed.
// Returns true if a token is consumed successfully, or false if the limit is hit.
func (rl *RateLimiter) Allow() bool {
	select {
	case <-rl.tokens:
		// Token available and successfully acquired
		return true
	default:
		// No tokens left in the bucket
		return false
	}
}

func main() {
	// Create a rate limiter: capacity of 5 tokens, refilling 1 token every 200ms
	limiter := NewRateLimiter(5, 200*time.Millisecond)

	fmt.Println("--- Simulating rapid requests ---")
	// Fire 10 requests rapidly (exceeds the 5-token burst limit)
	for i := 1; i <= 10; i++ {
		if limiter.Allow() {
			fmt.Printf("Request %2d: Allowed\n", i)
		} else {
			fmt.Printf("Request %2d: Denied (Rate limit exceeded)\n", i)
		}
		// Small delay between requests to watch behavior
		time.Sleep(200 * time.Millisecond)
	}

	// Wait long enough for the background goroutine to refill some tokens
	fmt.Println("\n--- Waiting 1 second for tokens to refill ---")
	time.Sleep(1 * time.Second)

	fmt.Println("--- Simulating requests after refill ---")
	for i := 11; i <= 15; i++ {
		if limiter.Allow() {
			fmt.Printf("Request %2d: Allowed\n", i)
		} else {
			fmt.Printf("Request %2d: Denied\n", i)
		}
	}
}
