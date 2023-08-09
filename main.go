package main

import (
	"github.com/symonk/learning-golang/constants"
	"github.com/symonk/learning-golang/forloop"
	"github.com/symonk/learning-golang/helloworld"
	"github.com/symonk/learning-golang/ifelse"
	"github.com/symonk/learning-golang/switches"
	"github.com/symonk/learning-golang/values"
	"github.com/symonk/learning-golang/variables"
)

func main() {
	/* Run all modules */
	helloworld.Run()
	values.Run()
	variables.Run()
	constants.Run()
	forloop.Run()
	ifelse.Run()
	switches.Run()
}
