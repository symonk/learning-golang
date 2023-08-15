package workerpools

import (
	"fmt"
	"time"
)

/*
Work groups are a mechanism to dispatch tasks off to multiple goroutines.
*/
func Run() {
	jobsToDo := 10
	tasks := make(chan int, jobsToDo)
	results := make(chan bool, jobsToDo)

	for i := 1; i <= 3; i++ {
		// Lets create 3 workers (we have 10 jobs to do but that's ok)
		// Keeping workers < jobs will simulate workers doing multiple
		go worker(i, tasks, results)
	}

	for j := 0; j < jobsToDo; j++ {
		// Push a new task onto the channel
		tasks <- j
	}
	// All tasks have been pushed on, close the channel as there is
	// Going to be no more tasks
	close(tasks)

	// Lets print results as and when they are completed
	for k := 0; k < jobsToDo; k++ {
		fmt.Println("Got a result", <-results)
	}

}

/*
The worker function will be spawned off in multiple goroutines concurrently.
These workers will receive work across a tasks channel and return their
results via sending to the results channel.  For now this simulates a long
running, possibly IO bound task by using time.Sleep to block the goroutine.
For now this is trivial boolean results, but you get the idea.
*/
func worker(id int, tasks <-chan int, results chan<- bool) {
	for task := range tasks {
		// Take a task and do it
		fmt.Println("Worker", id, "started job", task)
		time.Sleep(time.Second)
		fmt.Println("Worker", id, "finished job", task)
		// Push the result into the results channel
		results <- true
	}
}
