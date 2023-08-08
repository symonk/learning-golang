package main

import (
	"github.com/symonk/learning-golang/constants"
	"github.com/symonk/learning-golang/helloworld"
	"github.com/symonk/learning-golang/values"
	"github.com/symonk/learning-golang/variables"
)

func main() {
	/* Run all modules */
	helloworld.Run()
	values.Run()
	variables.Run()
	constants.Run()
}
