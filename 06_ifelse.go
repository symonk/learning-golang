package main

import "fmt"

func simpleIf(x bool) {
	if x {
		fmt.Println("True!")
		return
	}
}

func ifElse(val string) {
	if len(val) > 10 {
		fmt.Println("Longer than 10")
	} else {
		fmt.Println("Lesser than 10")
	}
}

func walrusAvailableInBranch() {
	if num := 9; num < 0 {
		fmt.Println("Less than.")
	} else if num == 9 {
		fmt.Println("Number was 9!")
	} else {
		fmt.Println("Greater than.")
	}

}

func main() {
	simpleIf(true)
	simpleIf(false)
	ifElse("HelloWorldExample")
	ifElse("short")
	walrusAvailableInBranch()
}
