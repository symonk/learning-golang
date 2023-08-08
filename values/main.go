package values

import (
	"fmt"
	"reflect"
)

// Basic examples of some of the primitive types in golang.
func Run() {
	var myString string = "myString"
	printValueAndType(myString)

	myInteger := 100
	printValueAndType(myInteger)

	myFloat := 200.45
	printValueAndType(myFloat)

	myBoolean := true
	printValueAndType(myBoolean)

}

func printValueAndType(element any) {
	fmt.Printf("element was type (%s) with value of '%v'\n", reflect.TypeOf(element), element)

}
