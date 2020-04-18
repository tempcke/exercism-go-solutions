package robotname

import (
	"errors"
	"math/rand"
	"time"
)

const (
	nameLetters   = 2
	nameLength    = 5
	possibleCount = 26 * 26 * 10 * 10 * 10
)

var usedNames = map[string]bool{}

// Robot with a unique name
type Robot struct {
	name string
}

// Reset the robots name
func (r *Robot) Reset() {
	r.name, _ = genName()
}

// Name of the robot
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		name, err := genName()
		if err != nil {
			return "", err
		}
		r.name = name
	}
	return r.name, nil
}

func genName() (string, error) {
	if len(usedNames) == possibleCount {
		return "", errors.New("possible names exausted")
	}
	strName := randName()
	for usedNames[strName] {
		strName = randName()
	}
	usedNames[strName] = true
	return strName, nil
}

func randName() string {
	rand.Seed(time.Now().UnixNano())
	name := make([]rune, nameLength)
	i := 0
	for ; i < nameLetters; i++ {
		name[i] = 'A' + int32(rand.Intn(26))
	}
	for ; i < nameLength; i++ {
		name[i] = '0' + int32(rand.Intn(10))
	}
	return string(name)
}
