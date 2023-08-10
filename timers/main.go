package timers

import (
	"fmt"
	"time"
)

/*
Timers are a great mechanism for achieving:
  - Running code at set intervals.
  - Running code at some point in the future.

`Tickers` are very similar to `Timers` so take a look at those aswell.
*/
func Run() {
	basicTimer()
	cancellingATimer()
}

func basicTimer() {
	/*
		Timers represent a single event in the future.  They are underpinned by
		an unbuffered channel.  Creating a new timer takes a duration after which
		it will fire.  The channel on a timer is set on the `C` attribute.
	*/
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("Timer started... we are blocking waiting for 2 seconds.")
	<-timer.C // wait for the timer channel to be unblocked
	fmt.Println("Timer fired after a 2 second delay")
	timer.Reset(2 * time.Second)
	fmt.Println("Reset the timer to use again!")
	<-timer.C
}

func cancellingATimer() {
	// Timers can be cancelled before they fire
	t := time.NewTimer(60 * time.Second)
	// Launch an async goroutine to wait on the timer and fire a task
	go func() {
		<-t.C
		fmt.Println("The timer finished!")
	}()
	t.Stop()
	fmt.Println("The timer was cancelled")

}
