package stringfunctions

import (
	"fmt"
	"strings"
)

func Run() {
	s := "foo"
	// Check if a string contains a substring
	print("Contains", strings.Contains(s, "oo"))
	// Check the count of a particular substring
	print("Count", strings.Count(s, "o"))
	// Check for prefix
	print("Prefixed", strings.HasPrefix(s, "fo"))
	// Check for suffix
	print("Suffixed", strings.HasSuffix(s, "o"))
	// Fetch the (first) index for a particular rune
	print("Index", strings.Index(s, "f"))
	// Join a sequence into a string by some separator
	print("Join", strings.Join([]string{"foo", "bar", "baz"}, "->"))
	// Repeat a string n number of times
	print("Repeat", strings.Repeat(s, 3))
	// Replace n occurrences of a substring in string
	print("Replace", strings.Replace(s, "oo", "ii", 1))
	// Split a string
	print("Split", strings.Split(s, "->"))
	// Convert to lower case
	print("Lower", strings.ToLower(s))
	// Convert to upper case
	print("Upper", strings.ToUpper(s))
}

func print(p string, result any) {
	fmt.Println(fmt.Sprintf("%s:\t%v", p, result))
}
