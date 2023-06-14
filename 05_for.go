package main

import "fmt"

const sep string = "====="

func singleConditionForLoop(until int) {
	i := 1
	for i <= until {
		fmt.Println(i)
		i += 1
	}
}

func classForLoop(until int) {

	for i := 1; i < until; i++ {
		fmt.Println(i)
	}
}

func infiniteLoop(until int) {
	i := 1
	for {
		fmt.Println("Looping forever...")
		if i > until {
			break
		}
		i++
	}
}

func continueForLoop(until int) {
	for n := 0; n <= until; n++ {
		if n == 3 {
			fmt.Println("Matched 3!")
			continue
		}
	}
}

func main() {
	singleConditionForLoop(5)
	classForLoop(5)
	infiniteLoop(10)
	continueForLoop(10)
}
