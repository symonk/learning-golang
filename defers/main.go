package defers

import (
	"fmt"
	"os"
)

// We previously touched on panics, defer is used to essentially
// guarantee clean up when functions exit (or fail) to avoid things
// like leaking resources etc, similar to pythons context managers.
func Run() {
	fileHandle := createFile("foo.txt")
	// Instantly defer the closing.
	defer closeFile(fileHandle)
	writeFile(fileHandle)
	fmt.Println("Finished writing data to the file", fileHandle.Name())

}

// Attempt to create a new file.
func createFile(p string) *os.File {
	file, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return file
}

// Close the file
func closeFile(f *os.File) {
	if err := f.Close(); err != nil {
		fmt.Println("Failure closing file", err)
		os.Exit(1)
	}
}

// Write data to the file
func writeFile(f *os.File) int {
	bytes, err := fmt.Fprintln(f, "writing into the file!")
	if err != nil {
		panic(err)
	}
	return bytes
}
