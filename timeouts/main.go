package timeouts

import (
	"fmt"
	"time"
)

/* Timeouts are important for programs that connect to external resources or that otehrwise
need to bound execution time.
*/

func Run() {
	afterTimeout()
	beforeTimeout()
}

func afterTimeout() {
	channelOne := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		channelOne <- "SentToOne"
	}()

	select {
	case result := <-channelOne:
		fmt.Println(result)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout 1!")
	}
}

func beforeTimeout() {
	channelTwo := make(chan int, 1)
	go func() {
		time.Sleep(1 * time.Second)
		channelTwo <- 1337
	}()
	select {
	case result := <-channelTwo:
		fmt.Println(result)
	case <-time.After(22 * time.Second):
		fmt.Println("We got a result before, this never runs.")
	}
}
