package strain

// Ints is an array of ints
type Ints []int

// Lists is an array of an array of ints
type Lists [][]int

// Strings is an array of strings
type Strings []string

// Keep takes a function used to decide which items to include in the returned set
func (c Ints) Keep(f func(int) bool) (bag Ints) {
	for _, item := range c {
		if f(item) {
			bag = append(bag, item)
		}
	}
	return
}

// Discard takes a function used to decide which items to exclude in the returned set
func (c Ints) Discard(f func(int) bool) Ints {
	return c.Keep(func(n int) bool { return !f(n) })
}

// Keep takes a function used to decide which items to include in the returned set
func (c Lists) Keep(f func([]int) bool) (bag Lists) {
	for _, item := range c {
		if f(item) {
			bag = append(bag, item)
		}
	}
	return
}

// Keep takes a function used to decide which items to include in the returned set
func (c Strings) Keep(f func(string) bool) (bag Strings) {
	for _, item := range c {
		if f(item) {
			bag = append(bag, item)
		}
	}
	return
}
