package regex

import (
	"fmt"
	"regexp"
)

// Go offers built in support for regular expressions
func Run() {
	fmt.Println(stringContainsAnyMatch("^hello.*", "hello world"))
}

// Check if a string has any matches of a given pattern.
func stringContainsAnyMatch(pattern string, str string) bool {
	match, err := regexp.MatchString(pattern, str)
	if err != nil {
		panic(err)
	}
	return match
}
