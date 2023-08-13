package main

import (
	"fmt"
	"reflect"
	"runtime"

	"github.com/symonk/learning-golang/arrays"
	"github.com/symonk/learning-golang/closures"
	"github.com/symonk/learning-golang/constants"
	"github.com/symonk/learning-golang/forloop"
	"github.com/symonk/learning-golang/functions"
	"github.com/symonk/learning-golang/helloworld"
	"github.com/symonk/learning-golang/ifelse"
	"github.com/symonk/learning-golang/maps"
	"github.com/symonk/learning-golang/pointers"
	"github.com/symonk/learning-golang/ranges"
	"github.com/symonk/learning-golang/recursion"
	"github.com/symonk/learning-golang/slices"
	"github.com/symonk/learning-golang/stringsrunes"
	"github.com/symonk/learning-golang/switches"
	"github.com/symonk/learning-golang/tickers"
	"github.com/symonk/learning-golang/timers"
	"github.com/symonk/learning-golang/values"
	"github.com/symonk/learning-golang/variables"
	"github.com/symonk/learning-golang/workergroups"
)

var (
	allFunctions = []func(){
		helloworld.Run,
		values.Run,
		variables.Run,
		constants.Run,
		forloop.Run,
		ifelse.Run,
		switches.Run,
		arrays.Run,
		slices.Run,
		maps.Run,
		ranges.Run,
		timers.Run,
		tickers.Run,
		workergroups.Run,
		functions.Run,
		variables.Run,
		closures.Run,
		recursion.Run,
		pointers.Run,
		stringsrunes.Run,
	}
)

func main() {
	/* Run all modules synchronously.  Modules have avoided builtins and are occassionally
	pluralised in their naming, i.e switches. */
	for _, callable := range allFunctions {
		fmt.Println("-----")
		fmt.Printf("Executing %s\n", runtime.FuncForPC(reflect.ValueOf(callable).Pointer()).Name())
		callable()
	}
}
