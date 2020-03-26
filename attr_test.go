package gom

import (
	"testing"
)

func NameTest(t *testing.T) {
	attr := createAttribute("ID")

	if attr.Name() == "id" {
		t.Fail()
	}
}
