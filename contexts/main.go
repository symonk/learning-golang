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
	deadlineCancellation()
}

func deadlineCancellation() {

	ctx, cancel := context.WithDeadlineCause(context.Background(), time.Now().Add(10*time.Second), errors.New("Deadline exceeded!"))
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context deadline firing...")
			panic(ctx.Err())
		default:
			fmt.Println("Still doing some heavy IO.")
			time.Sleep(time.Second)
		}
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
