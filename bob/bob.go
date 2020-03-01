// Package bob is a lackadaisical teenager with very limited responses
package bob

import "strings"

const (
	responseToSilence         = "Fine. Be that way!"
	responseToYellingQuestion = "Calm down, I know what I'm doing!"
	responseToYelling         = "Whoa, chill out!"
	responseToQuestion        = "Sure."
	response                  = "Whatever."
)

// Hey takes a remark and produces a lackadaisical response
func Hey(remark string) string {
	r := Remark{strings.TrimSpace(remark)}
	switch {
	case r.isSilence():
		return responseToSilence
	case r.isYelledQuestion():
		return responseToYellingQuestion
	case r.isYelling():
		return responseToYelling
	case r.isQuestion():
		return responseToQuestion
	default:
		return response
	}
}

// Remark is the comment made to bob
type Remark struct{ remark string }

func (r Remark) isYelledQuestion() bool {
	return r.isYelling() && r.isQuestion()
}

func (r Remark) isSilence() bool {
	return r.remark == ""
}

func (r Remark) isQuestion() bool {
	return strings.HasSuffix(r.remark, "?")
}

func (r Remark) isYelling() bool {
	matchesUpper := r.remark == strings.ToUpper(r.remark)
	matchesLower := r.remark == strings.ToLower(r.remark)
	return matchesUpper && !matchesLower
}
