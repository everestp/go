package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	mu        sync.Mutex
	count     int
	limit     int
	window    time.Duration
	resetTime time.Time
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:  limit,
		window: window,
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	now := time.Now()
	if now.After(rl.resetTime) {
		rl.resetTime = now.Add(rl.window)
		rl.count = 0

	}
	if rl.count < rl.limit {
		rl.count++
		return true
	}
	return false
}

func main() {
	rateLimiter := NewRateLimiter(2, 2*time.Second)

	for i := range 10 {
		if rateLimiter.Allow() {
			fmt.Println("Request allowed", i)
		} else {
			fmt.Println("Request denied")
		}
		time.Sleep(200 * time.Millisecond)
	}
}
