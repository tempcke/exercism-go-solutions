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
	c := make(chan FreqMap)
	defer close(c)

	m := FreqMap{}

	// wait for len(strings) processes to finish
	var wg sync.WaitGroup
	wg.Add(len(strings))

	// listen for new maps on chan
	// add them to our master map
	go func() {
		for freqMap := range c {
			for r, n := range freqMap {
				m[r] += n
			}
			wg.Done()
		}
	}()

	// for each string create a process to feed the FreqMap
	// into the channel then mark the process done
	for _, s := range strings {
		go func(s string) {
			c <- Frequency(s)
		}(s)
	}

	wg.Wait()
	return m
}
