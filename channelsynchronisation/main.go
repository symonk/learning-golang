package channelsynchronisation

import (
	"fmt"
	"time"
)

/*
We can use channels to synchronise execution across goroutines.
*/
func Run() {
	WaitForAGoroutineToFinish()
}

// It is common to use a boolean value to signal done, this is not advised
// using a struct{} is preferred as it saves the tiniest bit of memory and
// there is no need to send an actual value into the channel doing this
// way.
func WaitForAGoroutineToFinish() {
	done := make(chan struct{})
	go func() {
		fmt.Println("Doing some work..")
		// Simulate some long running work
		time.Sleep(5 * time.Second)
		fmt.Println("Finished the work!")
		defer close(done)
	}()
	// If we removed this, the program would exit instantly.
	<-done
}
