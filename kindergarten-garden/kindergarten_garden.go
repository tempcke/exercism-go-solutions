package kindergarten

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

var (
	rowCount         = 2
	cupsPerKidPerRow = 2
)

var plant = map[byte]string{
	'R': "radishes",
	'C': "clover",
	'G': "grass",
	'V': "violets",
}

// Garden represents the class garden
type Garden map[string][]string

// NewGarden constructs and returns a Garden
func NewGarden(diagram string, children []string) (*Garden, error) {
	g := &gardenBuilder{
		children: append([]string{}, children...),
		rows:     strings.Split(diagram, "\n")[1:], // diagram must start with newline
	}

	garden, err := g.build()

	return &garden, err
}

// Plants is used to query what plants each child has
func (g Garden) Plants(child string) ([]string, bool) {
	plants, ok := g[child]
	return plants, ok
}

// gardenBuilder exists simply to share state among methods
// while the Garden is being constructed and validated
type gardenBuilder struct {
	children []string
	rows     []string
}

func (g *gardenBuilder) validate() error {
	if err := g.validateChildren(); err != nil {
		return err
	}
	return g.validateRows()
}

func (g *gardenBuilder) validateRows() error {
	if len(g.rows) != rowCount {
		return errors.New("Invalid Diagram")
	}

	// validate rows are the correct length
	expectedLength := len(g.children) * cupsPerKidPerRow
	for _, row := range g.rows {
		if len(row) != expectedLength {
			return errors.New("Wrong number of cups")
		}

		// validate each cup in each row is valid
		for _, cup := range []byte(row) {
			if _, ok := plant[cup]; !ok {
				return fmt.Errorf("invalid cup code: %s", string(cup))
			}
		}
	}
	return nil
}

func (g *gardenBuilder) validateChildren() error {
	var prev string
	for _, name := range g.children {
		if name == prev {
			return errors.New("duplicate child names are not allowed")
		}
		prev = name
	}
	return nil
}

func (g *gardenBuilder) build() (Garden, error) {
	sort.Strings(g.children)

	if err := g.validate(); err != nil {
		return nil, err
	}
	return g.getGarden(), nil
}

func (g *gardenBuilder) getGarden() Garden {
	garden := make(map[string][]string, len(g.children))
	for n, name := range g.children {
		garden[name] = make([]string, 0, rowCount*cupsPerKidPerRow)

		for _, row := range g.rows {
			for j := 0; j < cupsPerKidPerRow; j++ {
				cup := row[cupsPerKidPerRow*n+j]
				garden[name] = append(garden[name], plant[cup])
			}
		}
	}
	return garden
}
