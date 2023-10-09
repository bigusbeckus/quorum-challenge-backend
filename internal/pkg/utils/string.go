package utils

import "regexp"

var nonAlphaNumericRegex = regexp.MustCompile("[^a-zA-Z0-9 ]+")

func RemoveNonAlphanumeric(s string) string {
	return nonAlphaNumericRegex.ReplaceAllString(s, " ")
}
