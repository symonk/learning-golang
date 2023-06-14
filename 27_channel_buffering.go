package main

import "fmt"

func main() {
	messages := make(chan string, 2)
	messages <- "foo"
	fmt.Println(<-messages)
	messages <- "bar"
	fmt.Println(<-messages)
}
