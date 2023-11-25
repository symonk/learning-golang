package signals

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Go offers first hand support for handling UNIX signals.
// We may want our programs to gracefully handle signals
// such as SIGTERM etc, here is how that can be done.
func Run() {
	handleSignal()

}

// Go handles signals by using a buffered channel
func handleSignal() {
	signals := make(chan os.Signal, 1)
	// Notify on sigterm and keyboard interrupt.
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	shouldExit := make(chan bool, 1)
	go func() {
		// Read from the signals channel if something appears.
		for {
			select {
			case sig := <-signals:
				fmt.Println(sig)
				shouldExit <- true
			case <-time.After(5 * time.Second):
				// Stop hanging indefinitely if user doesn't input an appropriate signal.
				fmt.Println("Timeout passed, exiting because no signal was sent!")
				shouldExit <- true
			}
		}
	}()
	// Block until the signal has been received.
	<-shouldExit
	fmt.Println("Gracefully exiting...")
}
