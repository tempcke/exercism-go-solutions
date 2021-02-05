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
	case len(a) == len(b) && hasSamePrefix(a, b):
		return Equal
	case isSubList(a, b):
		return SubList
	case isSubList(b, a):
		return SuperList
	}
	return UnEqual
}

func hasSamePrefix(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func isSubList(a, b []int) bool {
	return len(a) <= len(b) &&
		(hasSamePrefix(a, b) || isSubList(a, b[1:]))
}
