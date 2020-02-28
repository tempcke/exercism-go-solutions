package reverse

// Reverse an input string.  eg:Reverse("abc") returns "cba"
func Reverse(input string) string {
	inputRunes := []rune(input)
	n := len(inputRunes)
	outputRunes := make([]rune, n)
	for i := n - 1; i >= 0; i-- {
		outputRunes[n-i-1] = inputRunes[i]
	}
	return string(outputRunes)
}
