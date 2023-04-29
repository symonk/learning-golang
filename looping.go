package main

import (
	"fmt"
	"math/rand"
)

func main() {
	i := 1
	// A very basic fo loop
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// A for loop thats also initialised
	// & not incremented internally
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	options := []string{
		"One",
		"Two",
		"Three",
	}
	for {
		fmt.Println("looping forever!")
		random := options[rand.Intn(len(options))]
		if random == "Three" {
			fmt.Println("Hit the target, exiting the infinite for loop")
			break
		}
	}

	// continuing to the next iteration
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
