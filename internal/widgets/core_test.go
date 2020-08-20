package widgets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCore_Name(t *testing.T) {
	name := "core_widget"
	c := newCore(name)

	assert.Equal(t, c.name, name)
	assert.Equal(t, c.name, c.Name())
}

func TestCore_Parent(t *testing.T) {
	//parent := newMockLayout()
	//child :=
}
