package switches

import (
	"fmt"
	"time"
)

/*
Expressing conditionals across multiple branches
*/
func Run() {
	basicSwitch()
	switchGroupingCases()
	defaultSwitch()
}

func basicSwitch() {
	i := 0
	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("Two")
	}
}

func switchGroupingCases() {

	today := time.Now().Weekday()

	switch today {
	case time.Sunday, time.Saturday:
		fmt.Println("It is a weekend!")
	case time.Monday:
		fmt.Println("Monday")
	case time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("It is not monday or a weekend.")
	}

}

func defaultSwitch() {
	typeDetector := func(i any) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a boolean!")
		case int:
			fmt.Println("I'm an int!")
		case string:
			fmt.Println("I'm a string!")
		default:
			fmt.Printf("I didn't match any cases, unknown type %T\n", t)
		}
	}
	typeDetector(true)
	typeDetector(100)
	typeDetector("foo")
	typeDetector(nil)
}
