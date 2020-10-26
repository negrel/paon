package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExplode(t *testing.T) {
	splitChars := []string{"-", " ", "\t"}
	str := "Hello world, this-is an-hyphen"

	expected := []string{
		"Hello",
		"world,",
		"this",
		"is",
		"an",
		"hyphen",
	}

	assert.EqualValues(t, expected, Explode(str, splitChars...))
}

func TestWordWrap(t *testing.T) {
	str := "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
	expected := `Lorem Ipsum is simply dummy text of the
printing and typesetting industry. Lorem
Ipsum has been the industry's standard
dummy text ever since the 1500s, when an
unknown printer took a galley of type
and scrambled it to make a type specimen
book. It has survived not only five
centuries, but also the leap into
electronic typesetting, remaining
essentially unchanged. It was
popularised in the 1960s with the
release of Letraset sheets containing
Lorem Ipsum passages, and more recently
with desktop publishing software like
Aldus PageMaker including versions of
Lorem Ipsum.`

	assert.EqualValues(t, expected, WordWrap(str, 40))
}
