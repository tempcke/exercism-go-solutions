package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in all strings
// processing each string in parallel
func ConcurrentFrequency(texts []string) FreqMap {
	c := make(chan FreqMap)
	defer close(c)

	// call Frequency once for each text, pipe into chan
	for _, s := range texts {
		go func(s string) {
			c <- Frequency(s)
		}(s)
	}

	m := FreqMap{}

	// use texts just as a count for how many times to receive from chan
	for range texts {
		freqMap := <-c
		for r, n := range freqMap {
			m[r] += n
		}
	}
	return m
}
