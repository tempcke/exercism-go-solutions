// Package bob is a lackadaisical teenager with very limited responses
package bob

import "strings"

const (
	responseToEmpty           = "Fine. Be that way!"
	responseToYellingQuestion = "Calm down, I know what I'm doing!"
	responseToYelling         = "Whoa, chill out!"
	responseToQuestion        = "Sure."
	response                  = "Whatever."
)

// Hey takes a remark and produces a lackadaisical response
func Hey(remark string) string {
	trimmedRemark := strings.TrimSpace(remark)
	switch {
	case isEmpty(trimmedRemark):
		return responseToEmpty
	case isYellingQuestion(trimmedRemark):
		return responseToYellingQuestion
	case isYelling(trimmedRemark):
		return responseToYelling
	case isQuestion(trimmedRemark):
		return responseToQuestion
	default:
		return response
	}
}

func isYellingQuestion(trimmedRemark string) bool {
	return isYelling(trimmedRemark) && isQuestion(trimmedRemark)
}

func isEmpty(trimmedRemark string) bool {
	return trimmedRemark == ""
}

func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func isYelling(remark string) bool {
	return remark == strings.ToUpper(remark) && remark != strings.ToLower(remark)
}
