package channeldirections

import "fmt"

/*
When using channels as function parameters, it is possible to specify if a
channel is meant for retrieving or sending values to/from a channel.  This
helps catch more bugs and increases type-safety
*/
func Run() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "Passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

}

func ping(pings chan<- string, message string) {
	pings <- message
}

func pong(pings <-chan string, pongs chan<- string) {
	message := <-pings
	pongs <- message
}
