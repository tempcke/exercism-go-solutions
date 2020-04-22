package phonenumber

import (
	"errors"
	"fmt"
	"strings"
)

const format = "(%s) %s-%s"

type phNum struct {
	number string
}

func newPhNum(input string) (phNum, error) {
	var num strings.Builder

	for _, r := range input {
		if '0' <= r && r <= '9' {
			if num.Len() == 0 && r == '1' {
				continue
			}
			num.WriteRune(r)
		}
	}

	p := phNum{number: num.String()}
	return p, p.validate()
}

func (p phNum) validate() error {
	switch {
	case len(p.number) < 10 || len(p.number) > 11:
		return errors.New("Number must be 10 or 11 digits long")

	case len(p.number) == 11 && p.number[0] != '1':
		return errors.New("invalid when 11 digits does not start with a 1")

	case p.area()[0] == '0':
		return errors.New("area code may not start with zero")

	case p.exchange()[0] == '0' || p.exchange()[0] == '1':
		return errors.New("exchange code may not start with zero or one")
	}
	return nil
}

func (p phNum) area() string {
	return p.number[0:3]
}

func (p phNum) exchange() string {
	return p.number[3:6]
}

func (p phNum) subscriber() string {
	return p.number[6:]
}

// Number cleans up differently formated telephone numbers
// result is just 10 digits.  Ex: 6139950253
func Number(input string) (string, error) {
	p, err := newPhNum(input)

	if err != nil {
		return "", err
	}

	return p.number, nil
}

// AreaCode returns the area code of a phone number
func AreaCode(input string) (string, error) {
	p, err := newPhNum(input)

	if err != nil {
		return "", err
	}

	return p.area(), nil
}

// Format phone number in the format of (613) 995-0253
func Format(input string) (string, error) {
	p, err := newPhNum(input)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		format,
		p.area(),
		p.exchange(),
		p.subscriber(),
	), nil
}
