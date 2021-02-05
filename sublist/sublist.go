package sublist

// Relation between lists
type Relation string

// Relation values
const (
	Equal     Relation = "equal"
	SubList            = "sublist"
	SuperList          = "superlist"
	UnEqual            = "unequal"
)

// Sublist compares a and b to determine their Relation
// if a exists within b then a is a SubList of b
// if b exists within a then a is a SuperList of b
// if they are identical then a is Equal to b
// if a is not within b, nor b within a, then they are UnEqual
func Sublist(a, b []int) Relation {
	switch {
	case len(a) == len(b) && isEqual(a, b):
		return Equal
	case len(a) < len(b) && isSubList(a, b):
		return SubList
	case len(a) > len(b) && isSubList(b, a):
		return SuperList
	}
	return UnEqual
}

func isEqual(a, b []int) bool {
	for i := range a { // UnEqual unless identical
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func isSubList(a, b []int) bool {
	if len(a) == 0 {
		return true
	}

	for i := 0; len(a)+i <= len(b); i++ {
		if isMatch(a, b[i:]) {
			return true
		}
	}

	return false
}

func isMatch(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
