package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/negrel/debuggo/pkg/assert"
)

// Explode the given string using the given separators.
func Explode(s string, separators ...string) []string {
	return explode([]string{s}, separators...)
}

func explode(split []string, separators ...string) []string {
	if len(separators) == 0 {
		return split
	}

	var result []string
	for _, str := range split {
		result = append(result, strings.Split(str, separators[0])...)
	}

	return explode(result, separators[1:]...)
}

// WordWrap apply a wordwrap algorithm on the given string for the given
// width.
func WordWrap(s string, width int) string {
	assert.Greater(width, 0, "width must be greater than 0")

	// We add space so the regex can match the end of the string.
	s += " "

	re := fmt.Sprintf(`([^\n]{1,%v})\s`, width)
	regex := regexp.MustCompile(re)
	result := regex.ReplaceAllString(s, "$1\n")

	// We remove the previously added space
	return result[:len(result)-1]
}
