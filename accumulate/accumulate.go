// Package accumulate performs operation on each element of the collection
package accumulate

// Accumulate applies a function to each element
func Accumulate(strings []string, cb func(string) string) []string {
	result := make([]string, len(strings))
	for i, v := range strings {
		result[i] = cb(v)
	}
	return result
}
