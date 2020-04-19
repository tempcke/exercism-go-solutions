package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const possibleCount = 26 * 26 * 10 * 10 * 10

var usedNames = map[string]bool{}

// Robot with a unique name
type Robot struct {
	name string
}

// Name of the robot
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		if len(usedNames) == possibleCount {
			return "", errors.New("possible names exausted")
		}
		strName := randName()
		for usedNames[strName] {
			strName = randName()
		}
		usedNames[strName] = true
		r.name = strName
	}
	return r.name, nil
}

// Reset the robots name
func (r *Robot) Reset() {
	r.name = ""
}

func randName() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf(
		"%c%c%c%c%c",
		randLetter(),
		randLetter(),
		randDigit(),
		randDigit(),
		randDigit(),
	)
}

func randLetter() rune {
	return 'A' + int32(rand.Intn(26))
}
func randDigit() rune {
	return '0' + int32(rand.Intn(10))
}
