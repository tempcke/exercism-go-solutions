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

type teamMap map[string]team

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

		team0 := teams[tokens[0]]
		team1 := teams[tokens[1]]

		switch tokens[2] {
		case "win":
			recordMatch(&team0, &team1)
		case "loss":
			recordMatch(&team1, &team0)
		case "draw":
			recordDraw(&team0, &team1)
		default:
			return fmt.Errorf("Invalid line: %v", line)
		}

		teams[tokens[0]] = team0
		teams[tokens[1]] = team1
	}

	w.Write([]byte(header))
	for _, t := range teams.toSortedArray() {
		tallyRow(w, t)
	}

	return nil
}

func recordMatch(winner, looser *team) {
	winner.games++
	winner.wins++
	winner.points += winPoints

	looser.games++
	looser.losses++
}

func recordDraw(team1, team2 *team) {
	for _, t := range []*team{team1, team2} {
		t.games++
		t.draws++
		t.points += drawPoints
	}
}

func tallyRow(w io.Writer, t team) {
	fmt.Fprintf(w,
		rowFormat,
		t.name,
		t.games,
		t.wins,
		t.draws,
		t.losses,
		t.points,
	)
}

func (tm teamMap) toSortedArray() []team {
	teams := make([]team, 0, len(tm))
	for name, t := range tm {
		t.name = name
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
