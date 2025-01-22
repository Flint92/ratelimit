package token_bucket

import (
	"sync"
	"time"
)

type TokenBucket struct {
	sync.Mutex
	rate          int       // The speed at which tokens are generated, measured in tokens per second
	capacity      int       // The capacity of the bucket
	tokens        int       // The remaining tokens
	lastTimestamp time.Time // The last time of generate token
}

func NewTokenBucket(rate, capacity int) *TokenBucket {
	return &TokenBucket{
		rate:          rate,
		capacity:      capacity,
		tokens:        capacity,
		lastTimestamp: time.Now(),
	}
}

func (t *TokenBucket) Take() bool {
	t.Lock()
	defer t.Unlock()

	t.refill()
	if t.tokens > 0 {
		t.tokens--
		return true
	}

	return false
}

func (t *TokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(t.lastTimestamp)
	t.lastTimestamp = now
	newTokens := int(elapsed.Seconds()) * t.rate
	if newTokens > 0 {
		t.tokens += newTokens
		if t.tokens > t.capacity {
			t.tokens = t.capacity
		}
	}
}
