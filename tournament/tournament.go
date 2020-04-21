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
	rowFormat  = "%-30v | %2v | %2v | %2v | %2v | %2v\n"
)

var header = fmt.Sprintf(rowFormat, "Team", "MP", "W", "D", "L", "P")

type teamMap map[string]*team

type team struct {
	name   string
	games  int
	wins   int
	draws  int
	losses int
	points int
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

		team1 := teams.getTeam(tokens[0])
		team2 := teams.getTeam(tokens[1])

		switch tokens[2] {
		case "win":
			team1.recordWin()
			team2.recordLoss()
		case "loss":
			team1.recordLoss()
			team2.recordWin()
		case "draw":
			team1.recordDraw()
			team2.recordDraw()
		default:
			return fmt.Errorf("Invalid line: %v", line)
		}
	}

	w.Write([]byte(header))
	for _, t := range teams.toSortedArray() {
		tallyRow(w, t)
	}

	return nil
}

func tallyRow(w io.Writer, t *team) {
	w.Write([]byte(fmt.Sprintf(
		rowFormat,
		t.name,
		t.games,
		t.wins,
		t.draws,
		t.losses,
		t.points,
	)))
}

func (t *team) recordWin() {
	t.games++
	t.wins++
	t.points += winPoints
}

func (t *team) recordDraw() {
	t.games++
	t.draws++
	t.points += drawPoints
}

func (t *team) recordLoss() {
	t.games++
	t.losses++
}

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

func (tm *teamMap) getTeam(name string) *team {
	if t, ok := (*tm)[name]; ok {
		return t
	}
	t := &team{name: name}
	(*tm)[name] = t
	return t
}
