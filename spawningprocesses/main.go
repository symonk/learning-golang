package spawningprocesses

import (
	"fmt"
	"io"
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
	// piping basics
	commandWithPiping("grep", "hello")
}

func commandWithPiping(cmd string, args ...string) {
	grepCommand := exec.Command(cmd, args...)
	in, _ := grepCommand.StdinPipe()
	out, _ := grepCommand.StdoutPipe()
	grepCommand.Start()
	in.Write([]byte("hello grep\ngoodbye grep"))
	in.Close()
	grepBytes, _ := io.ReadAll(out)
	grepCommand.Wait()
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))
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
