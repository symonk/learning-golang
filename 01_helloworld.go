package main

import "fmt"

// https://gobyexample.com/hello-world
func HelloWorld(name string) string {
	message := fmt.Sprintf("Hello %v.", name)
	return message
}
