package widgets

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestWidgetList_Append(t *testing.T) {
	wl := NewWidgetList()
	w := NewNodeWidget("child")

	wl.Append(w)

	assert.Equal(t, 1, wl.Length())
	assert.Equal(t, w, wl.Get(0))
}

func TestWidgetList_GetIndexOf(t *testing.T) {
	wl := NewWidgetList()

	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("test-%v", i)
		w := NewNodeWidget(name)
		wl.Append(w)

		assert.Equal(t, i, wl.GetIndexOf(w))
	}
}

func TestWidgetList_Insert(t *testing.T) {
	wl := NewWidgetList()

	// Initialise the list
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("test-%v", i)
		w := NewNodeWidget(name)
		wl.Append(w)
	}

	// Insert at random index a widget
	for wl.Length() != 20 {
		index := rand.Intn(wl.Length())

		name := fmt.Sprintf("test-%v", index)
		w := NewNodeWidget(name)

		wl.Insert(w, index)

		assert.Equal(t, index, wl.GetIndexOf(w))
		assert.Equal(t, w, wl.Get(index))
	}
}

func TestWidgetList_Get(t *testing.T) {
	wl := NewWidgetList()
	child := NewNodeWidget("child")
	wl.Append(child)

	c := wl.Get(0)

	log.Println(child, c)

	assert.Equal(t, child, c)
}

func TestWidgetList_InsertBefore(t *testing.T) {
	wl := NewWidgetList()
	reference := NewNodeWidget("child")

	wl.Append(reference)

	w := NewNodeWidget("inserted_child")
	wl.InsertBefore(w, reference)

	assert.Equal(t, w, wl.Get(0))
	assert.Equal(t, reference, wl.Get(1))
}

func TestWidgetList_RemoveChild(t *testing.T) {
	wl := NewWidgetList()
	w := NewNodeWidget("child")

	wl.Append(w)

	assert.Equal(t, 1, wl.Length())

	wl.Remove(w)
	assert.Equal(t, 0, wl.Length())
	assert.Equal(t, NotFound, wl.GetIndexOf(w))
}
