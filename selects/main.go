package selects

import (
	"fmt"
	"time"
)

/*
Go select lets you wait on multiple channel operations.
Combining goroutines and channels with select is a powerful
feature.
*/
func Run() {
	// Create two unbuffered channels
	chan1 := make(chan string)
	chan2 := make(chan string)

	// Spawn two goroutines
	go func() {
		time.Sleep(1 * time.Second)
		chan1 <- "Pushed One."
	}()

	go func() {
		time.Sleep(2 * time.Second)
		chan2 <- "Pushed Two."
	}()

	for i := 0; i < 2; i++ {
		select {
		case message1 := <-chan1:
			fmt.Println("Received a value on channel 1", message1)
		case message2 := <-chan2:
			fmt.Println("Received a value on channel 2", message2)
		}
	}
}
