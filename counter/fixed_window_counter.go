package counter

import (
	"sync"
	"time"
)

type Counter struct {
	sync.Mutex
	limit       int
	windowSize  time.Duration
	requests    int
	windowStart time.Time
}

func NewCounter(limit int, windowSize time.Duration) *Counter {
	return &Counter{
		limit:       limit,
		windowSize:  windowSize,
		requests:    0,
		windowStart: time.Now(),
	}
}

func (c *Counter) Allow() bool {
	c.Lock()
	defer c.Unlock()
	now := time.Now()
	if now.Sub(c.windowStart) > c.windowSize {
		c.windowStart = now
		c.requests = 0
	}
	if c.requests < c.limit {
		c.requests++
		return true
	}
	return false
}
