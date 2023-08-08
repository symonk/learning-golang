package values

import (
	"github.com/symonk/learning-golang/utils"
)

// Basic examples of some of the primitive types in golang.
func Run() {
	var myString string = "myString"
	utils.PrintValueAndType(myString)

	myInteger := 100
	utils.PrintValueAndType(myInteger)

	myFloat := 200.45
	utils.PrintValueAndType(myFloat)

	myBoolean := true
	utils.PrintValueAndType(myBoolean)

}
