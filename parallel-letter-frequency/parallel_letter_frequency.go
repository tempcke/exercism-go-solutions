package letter

import (
	"sync"
)

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
func ConcurrentFrequency(strings []string) FreqMap {

	var wg sync.WaitGroup

	c := make(chan rune)

	m := FreqMap{}
	go func() {
		for r := range c {
			m[r]++
		}
	}()

	for _, s := range strings {
		wg.Add(1)
		go func(s string) {
			for _, r := range s {
				c <- r
			}
			wg.Done()
		}(s)
	}

	wg.Wait()
	close(c)

	return m

}
