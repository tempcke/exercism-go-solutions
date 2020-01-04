package twofer

import "fmt"

const DefaultName = "you"

func ShareWith(name string) string {
	msg := "One for %s, one for me."
	return fmt.Sprintf(msg, resolveName(name))
}

func resolveName(n string) string {
	if len(n) > 0 { return n }
	return DefaultName
}
