package regex

import (
	"fmt"
	"regexp"
)

// Go offers built in support for regular expressions
func Run() {
	fmt.Println(stringContainsAnyMatch("^hello.*", "hello world"))
	pattern := compiledPattern("^f")
	fmt.Println(pattern.MatchString("fo"))
	fmt.Println(pattern.MatchString("oof"))
	// matches `f` only.
	fmt.Println(pattern.FindString("foolproof"))
	pattern = compiledPattern("hello.*")
	// [3-14] index matching
	fmt.Println("index: ", pattern.FindStringIndex("ok hello there"))
	// Find information around full matches and sub matches within those full matches
	fmt.Println(pattern.FindStringSubmatch("ok hello world test"))
	// Replace all occurrences
	replaceAllOccurrences("hello[0-9]", "hello1, hello2, hello3, helloX, helloY", "<replaced>")
}

// Check if a string has any matches of a given pattern.
func stringContainsAnyMatch(pattern string, str string) bool {
	match, err := regexp.MatchString(pattern, str)
	if err != nil {
		panic(err)
	}
	return match
}

// Compiles are pattern for reuse. Panics if invalid pattern.
func compiledPattern(pattern string) *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

func replaceAllOccurrences(pattern string, s string, r string) {
	p := regexp.MustCompile(pattern)
	result := p.ReplaceAllString(s, r)
	fmt.Printf("Replaced %s with %s: %s\n", s, r, result)
}
