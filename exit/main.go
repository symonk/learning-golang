package exit

import (
	"fmt"
	"os"
)

// To exit a program immediately with a particular exit code
// `os.Exit` can be used.  However it will not run defers.
func Run() {
	exitWithoutDefer(3)
}

func exitWithoutDefer(code int) {
	defer func() { fmt.Println("I won't run!") }()
	os.Exit(code)
}
