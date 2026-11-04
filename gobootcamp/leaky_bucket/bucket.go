package main

import (
	"fmt"
	"sync"
	"time"
)

type LeakyBucket struct {
	capacity int
	leakRate time.Duration
	tokens   int
	lastleak time.Time
	mu       sync.Mutex
}

func NewLeakyBucket(capacity int, leakRate time.Duration) *LeakyBucket {
	return &LeakyBucket{
		capacity: capacity,
		leakRate: leakRate,
		tokens:   capacity,
		lastleak: time.Now(),
	}
}

func (lb *LeakyBucket) Allow() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	now := time.Now()
	elaspedTime := now.Sub(lb.lastleak)
	tokenstoAdd := int(elaspedTime / lb.leakRate)
	lb.tokens += tokenstoAdd

	if lb.tokens > lb.capacity {
		lb.tokens = lb.capacity
	}
	lb.lastleak = lb.lastleak.Add(time.Duration(tokenstoAdd) * lb.leakRate)
	// lb.lastleak = lb.lastleak.Add(elaspedTime) // wrong approach
	fmt.Printf("Token added %d , Token subtracted %d Total token %d \n", tokenstoAdd, 1, lb.tokens)
	fmt.Printf("last leak time : %v", lb.lastleak)
	if lb.tokens > 0 {
		lb.tokens--
		return true
	}
	return false

}

func main() {
	leakyBucketInst := NewLeakyBucket(5, 500*time.Millisecond)
	for range 10 {
		if leakyBucketInst.Allow() {
			fmt.Println("Request accepted")
		} else {
			fmt.Println(" XXXXX_-------Request rejected")
		}
		time.Sleep(time.Second)
	}
}
