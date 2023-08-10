package tickers

import (
	"fmt"
	"time"
)

/*
`Timers` are useful if you want to do something once in the future.
Tickers are for when you want to do something repeatidly at set
intervals, thing of `cron` on linux for example.  A ticker will
periodically tick until told to stop.
*/
func Run() {
	basicTicker()
}

func basicTicker() {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool) // Make a new channel to track if we are 'done'.

	go func() {
		for {
			select {
			case <-done:
				// We have finished, the ticker has been stopped.
				return
			case <-ticker.C:
				fmt.Println("Ticked again")
			}
		}
	}()
	time.Sleep(6 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Printf("Stopped the ticker, ticked %d times", len(done))
}
