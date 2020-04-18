package robotname

import (
	"math/rand"
	"time"
)

const (
	nameLetters = 2
	nameLength  = 5
)

var usedNames = map[string]bool{}

// Robot with a unique name
type Robot struct {
	name string
}

// Reset the robots name
func (r *Robot) Reset() {
	r.name = genName()
}

// Name of the robot
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.name = genName()
	}
	return r.name, nil
}

func genName() string {
	strName := randName()
	for usedNames[strName] {
		strName = randName()
	}
	usedNames[strName] = true
	return strName
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
