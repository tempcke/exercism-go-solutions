// Package twofer two for one. One for you and one for me.
package twofer

import "fmt"

const defaultName = "you"

// ShareWith Given a name, return a string with the message: One for X, one for me.
func ShareWith(name string) string {
	if len(name) == 0 {
		name = defaultName
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
