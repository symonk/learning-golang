package spawningprocesses

import (
	"fmt"
	"os/exec"
)

// Sometimes go programs need to spawn other processes
// Note: This examples are specific to a POSIX system
func Run() {
	// basic date output
	commandWithoutArgs("date")
	// date in UTC
	commandWithArgs("date", "-u")
	// error with args:
	commandWithArgs("date", "-x")
}

// A very basic command without any flags/args provided.
func commandWithoutArgs(cmd string) {
	command := exec.Command(cmd)
	output, err := command.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("Date was: ", string(output))
}

// A slightly more complicated command that provides options to the command
func commandWithArgs(cmd string, args ...string) {
	command := exec.Command(cmd, args...)
	output, err := command.Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("Failed executing: ", err)
		case *exec.ExitError:
			fmt.Println("command exit code: ", e.ExitCode())
		default:
			panic(err)
		}
	}
	fmt.Println("Date was: ", string(output))
}

func complexCommand(cmdstring string, args ...any) {

}
