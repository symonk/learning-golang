package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var ErrSentinelTimeout = errors.New("timed out")

// Often it is desirable to collect the output of various goroutines
// into one central error.  We can achieve this by using a waitgroup
// notifications channel and errors.Join
func main() {
	if err := run(); err != nil {
		fmt.Printf("We got an error %s", err.Error())
		if errors.Is(err, ErrSentinelTimeout) {
			fmt.Println("Sentinel error.")
		}
	}
}

func run() error {
	var wg sync.WaitGroup
	size := 10
	results := make(chan error, size)
	wg.Add(size)

	// spawn 10 goros; shoving their error results into the channel
	for i := 0; i < size; i++ {
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Second)
			results <- ErrSentinelTimeout
		}(i)
	}

	wg.Wait()
	close(results)
	combinedErrors := make([]error, size)
	for error := range results {
		combinedErrors = append(combinedErrors, error)
	}
	return errors.Join(combinedErrors...)
}
