package errgroups

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

/*


 */

var ErrNotOk = errors.New("not ok")

func Run() {
	errGroupsExample()
}

func errGroupsExample() {
	group := errgroup.Group{}
	for i := 0; i < 10; i++ {
		group.Go(func() error {
			fmt.Printf("Running task %d\n", i)
			time.Sleep(1 * time.Second)
			return ErrNotOk
		})
	}

	if err := group.Wait(); err != nil {
		fmt.Printf("Errors when running the tasks %v", err)
	} else {
		fmt.Printf("All tasks completed nicely.")
	}

}
