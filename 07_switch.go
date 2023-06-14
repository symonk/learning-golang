package main

import (
	"fmt"
)

func basicSwitch() {
	i := 1
	switch i {
	case 1:
		fmt.Println("one!")
	case 2:
		fmt.Println("never runs!")
	}
}

func switchManyCase() {
	i := 10
	switch i {
	case 8, 9, 10:
		fmt.Println("Ten!")
	}
}

func switchDefault() {
	var x string = "foo"
	switch x {
	case "nope":
		// nothing
	default:
		fmt.Println("X matched default!")
	}
}

func typeSwitch() {
	i := true
	myType := func(i interface{}) {
		switch myType := i.(type) {
		case bool:
			fmt.Println("I am a boolean!")
		default:
			fmt.Printf("I am type: %T\n", myType)
		}
	}
	myType(i)
	myType(100)
}

func main() {
	basicSwitch()
	switchManyCase()
	switchDefault()
	typeSwitch()
}
