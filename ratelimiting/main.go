package ratelimiting

import (
	"fmt"
	"time"
)

/*
Rate limiting is an important mechanism for controlling resource
utilisation and maintaining quality of service.  Go elegantly
supports rate limiting with goroutines, channels and tickers.
*/
func Run() {
	basicRateLimiting()
	burstyRateLimiting()
}

// Limit the channel reads to every 200 milliseconds.
func basicRateLimiting() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.NewTicker(1000 * time.Millisecond)

	for request := range requests {
		<-limiter.C
		fmt.Println("Throttled request", request, time.Now())
	}
}

// Allow short burts of request rates. Buffering the channel
// Allows upto 3 bursts of events.
func burstyRateLimiting() {
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.NewTicker(1000 * time.Millisecond).C {
			burstyLimiter <- t
		}
	}()

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	for request := range requests {
		<-burstyLimiter
		fmt.Println("request", request, time.Now())
	}

}
