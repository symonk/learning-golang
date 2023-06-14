package main

import "fmt"

func main() {
	// initialize a channel (unbuffered) and an anonymous function that passes a message into the channel
	// subsequently wait until a message is received by the channel and print it.
	messages := make(chan string)
	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}
