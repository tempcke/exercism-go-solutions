package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

const (
	delimiter  = ";"
	drawPoints = 1
	winPoints  = 3
)

type teamMap map[string]*team

func (tm teamMap) toSortedArray() []*team {
	teams := make([]*team, 0, len(tm))
	for _, t := range tm {
		teams = append(teams, t)
	}
	sort.Slice(teams, func(i, j int) bool {
		if teams[i].points == teams[j].points {
			return teams[i].name < teams[j].name
		}
		return teams[i].points > teams[j].points
	})
	return teams
}

type team struct {
	name   string
	games  int
	wins   int
	draws  int
	losses int
	points int
}

func (t *team) win() {
	t.games++
	t.wins++
	t.points += winPoints
}

func (t *team) draw() {
	t.games++
	t.draws++
	t.points += drawPoints
}

func (t *team) loss() {
	t.games++
	t.losses++
}

// Tally the results of a small football competition.
func Tally(r io.Reader, w io.Writer) error {
	var teams = make(teamMap, 2)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		tokens := strings.Split(line, delimiter)
		if len(tokens) != 3 {
			return fmt.Errorf("Invalid line: %v", line)
		}
		t1 := getTeam(teams, tokens[0])
		t2 := getTeam(teams, tokens[1])

		switch tokens[2] {
		case "win":
			t1.win()
			t2.loss()
		case "loss":
			t1.loss()
			t2.win()
		case "draw":
			t1.draw()
			t2.draw()
		default:
			return fmt.Errorf("Invalid line: %v", line)
		}
	}
	w.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))
	for _, t := range teams.toSortedArray() {
		teamStats := fmt.Sprintf(
			"%-30s | %2d | %2d | %2d | %2d | %2d\n",
			t.name,
			t.games,
			t.wins,
			t.draws,
			t.losses,
			t.points,
		)
		w.Write([]byte(teamStats))
	}
	return nil
}

func getTeam(teams teamMap, name string) *team {
	if t, ok := teams[name]; ok {
		return t
	}
	t := &team{name: name}
	teams[name] = t
	return t
}
