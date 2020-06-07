package core

import (
	"testing"
)

func TestName(t *testing.T) {
	name := "core"
	w := NewWidgetCore(name)

	if w.Name() != name {
		t.Fatalf("%v != %v, should be equal", w.Name(), name)
	}
}

func TestNextSibling(t *testing.T) {
	l := NewLayoutCore("parent")
	w1 := NewWidgetCore("child1")
	w2 := NewWidgetCore("child2")

	l.AppendChild(w1)
	l.AppendChild(w2)

	if nSib, ok := w1.NextSibling(); nSib != w2 || !ok {
		t.Fatalf("%+v != %+v, should be equal", nSib, w2)
	}

	if nSib, ok := w2.NextSibling(); nSib != nil || ok {
		t.Fatalf("%+v != %+v, should be equal", nSib, nil)
	}
}

func TestPreviousSibling(t *testing.T) {
	l := NewLayoutCore("parent")
	w1 := NewWidgetCore("child1")
	w2 := NewWidgetCore("child2")

	l.AppendChild(w1)
	l.AppendChild(w2)

	if pSib, ok := w2.PreviousSibling(); pSib != w1 || !ok {
		t.Fatalf("%+v != %+v, should be equal", pSib, w1)
	}

	if pSib, ok := w1.PreviousSibling(); pSib != nil || ok {
		t.Fatalf("%+v should be nil", pSib)
	}
}

func TestParent(t *testing.T) {
	l := NewLayoutCore("parent")
	w1 := NewWidgetCore("child1")
	w2 := NewWidgetCore("child2")

	l.AppendChild(w1)
	l.AppendChild(w2)

	if parent := w1.Parent(); parent != l {
		t.Fatalf("%+v != %+v, should be equal", parent, l)
	}

	if parent := w2.Parent(); parent != l {
		t.Fatalf("%+v != %+v, should be equal", parent, l)
	}

	if parent1, parent2 := w1.Parent(), w2.Parent(); parent1 != parent2 {
		t.Fatalf("%+v != %+v, should be equal", parent1, parent2)
	}

	if parent := l.Parent(); parent != nil {
		t.Fatalf("%+v should be nil", parent)
	}
}

func TestOwner(t *testing.T) {
	// TODO Widget method owner test
}

func TestSlot(t *testing.T) {
	l := NewLayoutCore("parent")
	w1 := NewWidgetCore("child1")
	w2 := NewWidgetCore("child2")

	l.AppendChild(w1)
	l.AppendChild(w2)

	if slot := w1.slot(); slot != 0 {
		t.Fatalf("%v should be 0", slot)
	}

	if slot := w2.slot(); slot != 1 {
		t.Fatalf("%v should be 1", slot)
	}

	l.DropChild(w1)

	if slot := w2.slot(); slot != 0 {
		t.Fatalf("%v should be 0", slot)
	}
}
