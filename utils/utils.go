package utils

import (
	"fmt"
	"reflect"
)

func PrintValueAndType(element any) {
	fmt.Printf("element was type (%s) with value of '%v'\n", reflect.TypeOf(element), element)

}
