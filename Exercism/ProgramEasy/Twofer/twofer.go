// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

import (
	"strings"
)

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	pattern := "One for you, one for me."
	if name != "" {
		return strings.Replace(pattern, "you", name, -1)
	}
	return pattern
}
