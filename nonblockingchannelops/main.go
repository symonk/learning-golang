package nonblockingchannelops

import (
	"fmt"
	"sync"
)

/*
	As we know, basic sends and receives on channels are blocking, however

incorporating `select` with a default calsue to implement non blocking sends,
receives and even non blocking multi-way selects is possible.
*/

func Run() {
	nonBlockingReceive()
	nonBlockingSend()
	multiwayNonBLockingSelect()
}

func nonBlockingReceive() {
	// An unbuffered channel, but reading does not block!
	messages := make(chan string)
	select {
	case message := <-messages:
		fmt.Println(message)
	default:
		fmt.Println("No new messages on the messages channel")
	}
}

func nonBlockingSend() {
	// Again an unbuffered channel, but sending does not block after 1 value.
	channel := make(chan int)
	// We cannot send this value in, as the channel is unbuffered and there
	// is no receiver
	val := 100
	select {
	case channel <- val:
		fmt.Println("Sent value", val)
	default:
		fmt.Println("No messages could be sent.")
	}
}

func multiwayNonBLockingSelect() {
	// We can use multiple cases above the default clause to
	// implement a multi way non blocking select.  Here we will
	// attempt non blocking receives on multiple channels.
	channel := make(chan int, 1)
	signals := make(chan bool, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		channel <- 100
		signals <- true
		wg.Done()

	}()
	wg.Wait()
	for {
		select {
		case retrieved := <-channel:
			fmt.Println("Retrieved a value", retrieved)
		case sig := <-signals:
			fmt.Println("Received a signal", sig)
		default:
			fmt.Println("No activity across both channels, exiting.")
			return
		}
	}

}
