package wordy

import (
	"strconv"
	"strings"
)

// Answer Parse and evaluate simple math word problems returning the answer as an integer
func Answer(question string) (answer int, ok bool) {
	q := strings.ToLower(question)
	q = strings.Replace(q, " by ", " ", -1)
	q = strings.Trim(q, "?")
	q = strings.TrimPrefix(q, "what is")
	fields := strings.Fields(q)

	if len(fields) == 0 {
		return 0, false
	}

	a, ok := parseInt(fields[0])

	if !ok {
		return 0, false
	}

	if len(fields) == 1 {
		return a, true
	}

	return eval(a, fields[1:])
}

func parseInt(input string) (n int, ok bool) {
	n, err := strconv.Atoi(input)
	if err != nil {
		return
	}
	return n, true
}

func eval(a int, fields []string) (answer int, ok bool) {
	if len(fields) == 0 {
		return a, true
	}

	// if a is an int and there is only one tok
	// it is invalid regardless if it is an operation or an int
	if len(fields) == 1 {
		return 0, false
	}

	b, ok := parseInt(fields[1])
	if !ok {
		return 0, false
	}

	switch fields[0] {
	case "plus":
		return eval(a+b, remainingFields(fields))
	case "minus":
		return eval(a-b, remainingFields(fields))
	case "multiplied":
		return eval(a*b, remainingFields(fields))
	case "divided":
		return eval(a/b, remainingFields(fields))
	}
	return 0, false
}

func remainingFields(toks []string) []string {
	if len(toks) == 2 {
		return []string{}
	}
	return toks[2:]
}
