package main

import (
	"fmt"
	"time"
)

/*
Similar to maps, channels are a reference type are are allocated with `make`.  If an optional
integer is provided, the channel is considered a buffered channel.  Channels without a buffer
set are known as synchronous, or unbuffered channels.

Channels combine communicate with synchronization, guaranteeing that multiple goroutines are
in a known state.
*/

func longRunningFunction(delay int) int {
	fmt.Printf("Sleeping for %d seconds.", delay)
	time.Sleep(time.Duration(delay) * time.Second)
	return delay
}

func channelBlockingOnGoroutines() int {
	theChannel := make(chan int)
	go func() {
		theChannel <- longRunningFunction(3)
	}()
	return <-theChannel

}

func main() {
	// Build an unbuffered channel of integers
	basicChannel := make(chan int)
	// Build a unbuffered/sync channel of strings
	// This is the same under the hood as `basicChannel` above.
	basicSyncChannel := make(chan string, 0)
	// Build a buffered channel of size 10 of strings
	basicBufferedChannel := make(chan string, 10)

	// close the channels
	close(basicChannel)
	close(basicSyncChannel)
	close(basicBufferedChannel)

	channelBlockingOnGoroutines()

}
