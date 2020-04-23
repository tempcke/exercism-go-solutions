package phonenumber

import (
	"errors"
	"fmt"
	"strings"
)

const format = "(%s) %s-%s"

// Number cleans up differently formated telephone numbers
// result is just 10 digits.  Ex: 6139950253
func Number(input string) (string, error) {

	var num strings.Builder
	for _, r := range input {
		if '0' <= r && r <= '9' {
			num.WriteRune(r)
		}
	}

	p := num.String()
	if len(p) == 11 && p[0] == '1' {
		p = p[1:]
	}

	switch {
	case len(p) != 10:
		return "", errors.New("Number must be 10, or 11 if it starts with 1")
	case p[0] <= '1':
		return "", errors.New("area code may not start with zero or one")
	case p[3] <= '1':
		return "", errors.New("exchange code may not start with zero or one")
	}

	return p, nil
}

// AreaCode returns the area code of a phone number
func AreaCode(input string) (string, error) {
	p, err := Number(input)

	if err != nil {
		return "", err
	}

	return p[0:3], nil
}

// Format phone number in the format of (613) 995-0253
func Format(input string) (string, error) {
	p, err := Number(input)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf(format, p[0:3], p[3:6], p[6:]), nil
}
