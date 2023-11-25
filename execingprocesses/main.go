package execingprocesses

import (
	"os"
	"os/exec"
	"syscall"
)

// Similarly to spawning a subprocess, exec
// can replace the go process with another one
// As always, taking user input into exec directly
// is a serious security flaw and should be used with caution.
func Run() {
	listFilesAndDirectories()
}

func listFilesAndDirectories() {
	executable, err := exec.LookPath("ls")
	if err != nil {
		// POSIX only here again.
		panic(err)
	}
	args := []string{"ls", "-a", "-l", "-h"}
	env := os.Environ()
	if err := syscall.Exec(executable, args, env); err != nil {
		panic(err)
	}

}
