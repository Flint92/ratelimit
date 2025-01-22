package leaky_bucket

import (
	"sync"
	"time"
)

type LeakyBucket struct {
	sync.Mutex
	rate          int
	capacity      int
	water         int
	lastTimestamp time.Time
}

func NewLeakyBucket(rate int, capacity int) *LeakyBucket {
	return &LeakyBucket{
		rate:          rate,
		capacity:      capacity,
		water:         0,
		lastTimestamp: time.Now(),
	}
}

func (lb *LeakyBucket) Allow() bool {
	lb.Lock()
	defer lb.Unlock()

	lb.drain()
	if lb.water < lb.capacity {
		lb.water++
		return true
	}

	return false
}

func (lb *LeakyBucket) drain() {
	now := time.Now()
	elapsed := now.Sub(lb.lastTimestamp)
	lb.lastTimestamp = now
	drained := int(elapsed.Seconds()) * lb.rate
	if drained > 0 {
		lb.water -= drained
		if lb.water < 0 {
			lb.water = 0
		}
	}
}
