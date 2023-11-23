package contexts

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// Contexts are useful for handling and controlling cancellation
// They also allow carrying request-scoped value across boundaries
// Note: CancelFuncs should always be closed to avoid leakages.
func Run() {
	// cancelling because of a deadline.
	if err := deadlineCancellation(); err != nil {
		fmt.Println("Received a deadline error: ", err.Error())
	}
	// cancelling because of a direct cancel
	if err := contextCancel(); err != nil {
		fmt.Println("Received an explicit context cancel: ", err.Error())
	}

	// cancelling because of a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := contextTimeout(ctx); err != nil {
		fmt.Println("Received a cancellation error: ", err.Error())
	}
}

func deadlineCancellation() error {

	ctx, cancel := context.WithDeadlineCause(context.Background(), time.Now().Add(3*time.Second), errors.New("Deadline exceeded!"))
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			fmt.Println("No context deadline yet, will keep doing IO")
			time.Sleep(time.Second)
		}
	}
}

func contextTimeout(ctx context.Context) error {
	// Block the channel with a simple goro.
	neverReady := make(chan struct{})
	select {
	case <-neverReady:
		//
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

func contextCancel() error {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(2 * time.Second)
		// force cancel!
		cancel()
	}()
	select {
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Do not store contexts for reuse in a struct
type DoNotDoThis struct {
	ctx context.Context
}

// Users have no control/scope to handle deadlines etc.
func (d *DoNotDoThis) Fetch() {
	fmt.Println("Performing some IO bound fetch.")
}

// Improved Version:
type DoThis struct {
}

func (d *DoThis) Fetch(ctx context.Context) {
	fmt.Println("Performing some IO bound fetch, but its independently cancellable etc.")
}
