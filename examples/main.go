package main

import (
	"fmt"
	fwc "github.com/flint92/ratelimit/counter"
	lb "github.com/flint92/ratelimit/leaky_bucket"
	tb "github.com/flint92/ratelimit/token_bucket"
	"sync"
	"time"
)

func main() {
	testTokenBucket()
	testLeakyBucket()
	testFixedWindowCounter()
}

func testFixedWindowCounter() {
	fmt.Println("============= test fixed window counter ===============")

	l := fwc.NewCounter(10, time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			if l.Allow() {
				fmt.Println("Request", num, "allowed")
			} else {
				fmt.Println("Request", num, "rejected")
			}
			time.Sleep(10 * time.Millisecond) // 模拟请求间隔
		}(i)
	}

	wg.Wait()
}

func testLeakyBucket() {
	fmt.Println("============= test leaky bucket ===============")

	l := lb.NewLeakyBucket(10, 10)

	var wg sync.WaitGroup
	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			if l.Allow() {
				fmt.Println("Request", num, "allowed")
			} else {
				fmt.Println("Request", num, "rejected")
			}
			time.Sleep(10 * time.Millisecond) // 模拟请求间隔
		}(i)
	}

	wg.Wait()
}

func testTokenBucket() {
	fmt.Println("============= test token bucket ===============")

	t := tb.NewTokenBucket(10, 10)

	var wg sync.WaitGroup
	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			if t.Take() {
				fmt.Println("Request", num, "allowed")
			} else {
				fmt.Println("Request", num, "rejected")
			}
			time.Sleep(10 * time.Millisecond) // 模拟请求间隔
		}(i)
	}

	wg.Wait()
}
