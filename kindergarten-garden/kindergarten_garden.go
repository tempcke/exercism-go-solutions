package kindergarten

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// Garden represents the class garden
type Garden struct {
	children []string
	plants   map[string][]byte
}

var plant = map[byte]string{
	'R': "radishes",
	'C': "clover",
	'G': "grass",
	'V': "violets",
}

// NewGarden constructs and returns a Garden
func NewGarden(diagram string, children []string) (*Garden, error) {
	g := new(Garden)
	g.children = append([]string{}, children...)
	sort.Strings(g.children)

	rows := strings.Split(diagram, "\n")
	if len(rows) != 3 || len(rows[1]) != len(children)*2 {
		return nil, errors.New("Wrong number of plants in diagram")
	}
	g.plants = make(map[string][]byte, len(children))

	for i, name := range g.children {
		if _, ok := g.plants[name]; ok {
			return nil, fmt.Errorf("duplicate name: %s", name)
		}
		cups := []byte{
			rows[1][i*2], rows[1][i*2+1],
			rows[2][i*2], rows[2][i*2+1],
		}
		for _, c := range cups {
			if _, ok := plant[c]; !ok {
				return nil, fmt.Errorf("invalid cup code: %s", string(c))
			}
		}
		g.plants[name] = cups
	}

	return g, nil
}

// Plants is used to query what plants each child has
func (g *Garden) Plants(child string) ([]string, bool) {
	bytes, ok := g.plants[child]
	if !ok {
		return nil, ok
	}

	plants := make([]string, len(bytes))
	for i, b := range bytes {
		plants[i] = plant[b]
	}

	return plants, ok
}
