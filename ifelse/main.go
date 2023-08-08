package ifelse

import "fmt"

/*
Basic branched logic in golang for the if, else, else if idioms
Switch exists but is a separate package here.
*/

func Run() {
	basicIf()
	basicIfElse()
	basicElseIfElse()
}

func basicIf() {
	if true {
		fmt.Println("It's true.")
	}
}

func basicIfElse() {
	predicate := true
	if !predicate {
		fmt.Println("This will never run.")
	} else {
		fmt.Println("Predicate was actually true.")
	}
}

func basicElseIfElse() {
	number := 100
	if number < 100 {
		panic("This can never run!")
	} else if number == 0 {
		panic("This can also never run!")
	} else {
		fmt.Println("Number is 100, we finally exit.")
	}
}
