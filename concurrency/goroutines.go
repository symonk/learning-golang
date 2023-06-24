package main

import (
	"fmt"
	"time"
)

/*
Goroutines are functions executing concurrently in the same address space space
as other goroutines.  They are extremely lightweight and cheap to use.

Goroutines are multiplexed into multiple operating system threads, if one should
block while waiting for IO or examples, others will continue to run.

Goroutines are not blocking obviously, so when calling them it is important
(if required) to handle managing waiting accordingly until things are done etc.
*/

func simpleFunction() {
	fmt.Println("Sleeping...")
	time.Sleep(time.Second * 5)
}

func main() {
	//  Invoke the function normally
	simpleFunction()

	// run the function concurrently; this is not blocking so the program can exit here
	// prior to the function returning.
	go simpleFunction()

}
