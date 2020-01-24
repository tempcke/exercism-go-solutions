package romannumerals

import "errors"

var m = []struct {
	r string
	v int
}{

	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

// ToRomanNumeral converts an int to a roman numeral string
func ToRomanNumeral(number int) (string, error) {
	if !canConvert(number) {
		return "", errors.New("number must be greater than zero and less than or equal to 3000")
	}
	return int2Rom(number), nil
}

func int2Rom(number int) string {
	var roman string
	n := number
	for n > 0 {
		for _, rmap := range m {
			if rmap.v <= n {
				roman += rmap.r
				n -= rmap.v
				break
			}
		}
	}
	return roman
}

func canConvert(n int) bool {
	if n <= 0 {
		return false
	}
	if n > 3000 {
		return false
	}
	return true
}
