package wordy

import (
	"strconv"
	"strings"
)

var strReplaceMap = map[string]string{
	"plus":          "+",
	"minus":         "-",
	"multiplied by": "*",
	"divided by":    "/",
	"raised to the": "^",
	"st power":      "",
	"nd power":      "",
	"rd power":      "",
	"th power":      "",
}

var ops = map[string]func(int, int) int{
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"*": func(a, b int) int { return a * b },
	"/": func(a, b int) int { return a / b },
	"^": func(a, b int) int {
		n := a
		for i := 0; i < b-1; i++ {
			n *= a
		}
		return n
	},
}

// Answer Parse and evaluate simple math word problems returning the answer as an integer
func Answer(question string) (answer int, ok bool) {
	// format check
	if !strings.HasPrefix(question, "What is ") || !strings.HasSuffix(question, "?") {
		return 0, false
	}

	// trim prefix and sufix
	q := question[8 : len(question)-1]

	// single char operators
	for substr, replacement := range strReplaceMap {
		q = strings.ReplaceAll(q, substr, replacement)
	}

	// split by whitespace
	fields := strings.Fields(q)

	// expression must be an odd number of ints and operations
	// becuase it must always start and end with a number
	if len(fields)%2 == 0 {
		return 0, false
	}

	// first int
	a, ok := parseInt(fields[0])
	if !ok {
		return 0, false
	}

	// expresion was just one number like: what is 5?
	if len(fields) == 1 {
		return a, true
	}

	// iterate next 2 fields at a time, operation and int
	for i := 1; i < len(fields)-1; i += 2 {
		op, ok := ops[fields[i]]
		if !ok {
			return 0, false
		}
		b, ok := parseInt(fields[i+1])
		if !ok {
			return 0, false
		}
		a = op(a, b)
	}

	return a, true
}

func parseInt(input string) (n int, ok bool) {
	n, err := strconv.Atoi(input)
	if err != nil {
		return 0, false
	}
	return n, true
}
