package closingchannels

import "fmt"

/*
	Closing a channel indicates that no more values will be sent on

it, receivers can check if a channel is closed to know themselves
when it may be sensible to exit.
*/
func Run() {
	closingChannel()
}

func closingChannel() {
	// Establish two channels, one for processable work
	// One to signal when things are done.
	workToDo := make(chan int, 3)
	workDone := make(chan struct{})

	// Spawn a goroutine that indefinitely waits for jobs.
	go func() {
		for {
			// Here we use the comma idiom, to see if thee channel is returning
			// falsy exhausted values, or there was actual values on it.
			task, more := <-workToDo
			if more {
				fmt.Println("Received an actual job", task)
			} else {
				fmt.Println("All jobs have been received")
				close(workDone)
				return
			}

		}
	}()

	for i := 0; i < 10; i++ {
		// Spawn off some work to do.
		// The receiver is ready to wait.
		workToDo <- i
		fmt.Println("Sent a task", i)
	}
	// Close the work to be done channel to signal to the
	// receiving goroutine it is finished.
	close(workToDo)
	// Wait for the `done` signal to be dispatched, so we can exit.
	<-workDone
}
