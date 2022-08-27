package utils

import (
	"regexp"
	"strings"
)

func SplitStringsByEmptyNewline(str string) []string {
	// Fix windows returns
	//nolint:gocritic
	str = strings.Replace(str, "\r", "\n", -1)

	return regexp.MustCompile(`\n\s*\n`).Split(str, -1)
}
