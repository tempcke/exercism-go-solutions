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
	if r.name != "" {
		return r.name, nil
	}

	if len(usedNames) == possibleCount {
		return "", errors.New("possible names exhausted")
	}

	r.name = randName()
	for usedNames[r.name] {
		r.name = randName()
	}
	usedNames[r.name] = true
	return r.name, nil
}

// Reset the robots name
func (r *Robot) Reset() {
	r.name = ""
}

func randName() string {
	return fmt.Sprintf(
		"%s%s%03d",
		string(rand.Intn(26)+'A'),
		string(rand.Intn(26)+'A'),
		rand.Intn(1000),
	)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
