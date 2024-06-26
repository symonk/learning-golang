package main

import (
	"flag"
	"fmt"

	"github.com/symonk/learning-golang/arrays"
	"github.com/symonk/learning-golang/atomiccounters"
	"github.com/symonk/learning-golang/bufferedchannels"
	"github.com/symonk/learning-golang/channeldirections"
	"github.com/symonk/learning-golang/channels"
	"github.com/symonk/learning-golang/channelsynchronisation"
	"github.com/symonk/learning-golang/closingchannels"
	"github.com/symonk/learning-golang/closures"
	"github.com/symonk/learning-golang/constants"
	"github.com/symonk/learning-golang/contexts"
	"github.com/symonk/learning-golang/defers"
	"github.com/symonk/learning-golang/errgroups"
	"github.com/symonk/learning-golang/errors"
	"github.com/symonk/learning-golang/execingprocesses"
	"github.com/symonk/learning-golang/exit"
	"github.com/symonk/learning-golang/forloop"
	"github.com/symonk/learning-golang/functions"
	"github.com/symonk/learning-golang/generics"
	"github.com/symonk/learning-golang/goroutines"
	"github.com/symonk/learning-golang/helloworld"
	"github.com/symonk/learning-golang/ifelse"
	"github.com/symonk/learning-golang/interfaces"
	"github.com/symonk/learning-golang/maps"
	"github.com/symonk/learning-golang/methods"
	"github.com/symonk/learning-golang/mutexes"
	"github.com/symonk/learning-golang/nonblockingchannelops"
	"github.com/symonk/learning-golang/panics"
	"github.com/symonk/learning-golang/pointers"
	"github.com/symonk/learning-golang/rangeoverchannels"
	"github.com/symonk/learning-golang/ranges"
	"github.com/symonk/learning-golang/ratelimiting"
	"github.com/symonk/learning-golang/recover"
	"github.com/symonk/learning-golang/recursion"
	"github.com/symonk/learning-golang/regex"
	"github.com/symonk/learning-golang/selects"
	"github.com/symonk/learning-golang/signals"
	"github.com/symonk/learning-golang/slices"
	"github.com/symonk/learning-golang/sorting"
	"github.com/symonk/learning-golang/sortingbyfunctions"
	"github.com/symonk/learning-golang/spawningprocesses"
	"github.com/symonk/learning-golang/statefulgoroutines"
	"github.com/symonk/learning-golang/stringformatting"
	"github.com/symonk/learning-golang/stringfunctions"
	"github.com/symonk/learning-golang/stringsrunes"
	"github.com/symonk/learning-golang/structembedding"
	"github.com/symonk/learning-golang/structs"
	"github.com/symonk/learning-golang/switches"
	"github.com/symonk/learning-golang/texttemplates"
	"github.com/symonk/learning-golang/tickers"
	"github.com/symonk/learning-golang/timeouts"
	"github.com/symonk/learning-golang/timers"
	"github.com/symonk/learning-golang/values"
	"github.com/symonk/learning-golang/variables"
	"github.com/symonk/learning-golang/waitgroups"
	"github.com/symonk/learning-golang/workerpools"
)

func main() {
	/* Run all modules synchronously.  Modules have avoided builtins and are occassionally
	pluralised in their naming, i.e switches. */
	module := flag.String("module", "all", "The name of the folder to run and exit, defaults to run all.")
	flag.Parse()

	options := buildMap()
	if *module == "all" {
		for _, value := range options {
			fmt.Println("-----")
			fmt.Printf("Executing %s\n", *module)
			value()

		}
	} else {
		callable, ok := options[*module]
		if !ok {
			panic("could not find a module named: `" + *module + "` Are you sure the folder name is correct?")
		}
		fmt.Printf("Executing %s\n", *module)
		callable()
	}
}

func buildMap() map[string]func() {
	fnMap := make(map[string]func())
	fnMap["helloworld"] = helloworld.Run
	fnMap["values"] = values.Run
	fnMap["variables"] = variables.Run
	fnMap["constants"] = constants.Run
	fnMap["forloop"] = forloop.Run
	fnMap["ifelse"] = ifelse.Run
	fnMap["switches"] = switches.Run
	fnMap["arrays"] = arrays.Run
	fnMap["slices"] = slices.Run
	fnMap["maps"] = maps.Run
	fnMap["ranges"] = ranges.Run
	fnMap["timers"] = timers.Run
	fnMap["tickers"] = tickers.Run
	fnMap["workerpools"] = workerpools.Run
	fnMap["functions"] = functions.Run
	fnMap["closures"] = closures.Run
	fnMap["recursion"] = recursion.Run
	fnMap["pointers"] = pointers.Run
	fnMap["stringsrunes"] = stringsrunes.Run
	fnMap["structs"] = structs.Run
	fnMap["methods"] = methods.Run
	fnMap["interfaces"] = interfaces.Run
	fnMap["structembedding"] = structembedding.Run
	fnMap["generics"] = generics.Run
	fnMap["errors"] = errors.Run
	fnMap["goroutines"] = goroutines.Run
	fnMap["channels"] = channels.Run
	fnMap["bufferedchannels"] = bufferedchannels.Run
	fnMap["channelsynchronisation"] = channelsynchronisation.Run
	fnMap["channeldirections"] = channeldirections.Run
	fnMap["selects"] = selects.Run
	fnMap["timeouts"] = timeouts.Run
	fnMap["nonblockingchannelops"] = nonblockingchannelops.Run
	fnMap["closingchannels"] = closingchannels.Run
	fnMap["rangeoverchannels"] = rangeoverchannels.Run
	fnMap["waitgroups"] = waitgroups.Run
	fnMap["ratelimiting"] = ratelimiting.Run
	fnMap["atomiccounters"] = atomiccounters.Run
	fnMap["mutexes"] = mutexes.Run
	fnMap["statefulgoroutines"] = statefulgoroutines.Run
	fnMap["sorting"] = sorting.Run
	fnMap["sortingbyfunctions"] = sortingbyfunctions.Run
	fnMap["panics"] = panics.Run
	fnMap["defer"] = defers.Run
	fnMap["recover"] = recover.Run
	fnMap["stringfunctions"] = stringfunctions.Run
	fnMap["stringformatting"] = stringformatting.Run
	fnMap["texttemplates"] = texttemplates.Run
	fnMap["regexp"] = regex.Run
	//
	fnMap["contexts"] = contexts.Run
	fnMap["spawningprocesses"] = spawningprocesses.Run
	fnMap["execingprocesses"] = execingprocesses.Run
	fnMap["signals"] = signals.Run
	fnMap["exit"] = exit.Run
	fnMap["errgroups"] = errgroups.Run
	return fnMap
}
