package draw

import (
	"github.com/negrel/paon/internal/geometry"
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randInt(start, end int) int {
	n := int64(rand.Int())
	scale := math.MaxInt64 / (end - start)

	return start + (int(n) / scale)
}

func TestMakeCanvas(t *testing.T) {
	w := rand.Int()
	h := rand.Int()
	canvas := MakeCanvas(geometry.Rect(0, 0, w, h))

	assert.Equal(t, canvas.Bounds.Width(), w)
	assert.Equal(t, canvas.Bounds.Height(), h)
	assert.NotNil(t, canvas.cellGrid)
}

func TestCanvas_Get_NoBounds(t *testing.T) {
	canvas := MakeCanvas(geometry.Rectangle{})

	assert.Nil(t, canvas.Get(geometry.Pt(0, 0)))
}

func TestCanvas_Get_OutsideBounds(t *testing.T) {
	w := rand.Int()
	h := rand.Int()
	canvas := MakeCanvas(geometry.Rect(0, 0, w, h))

	assert.Nil(t, canvas.Get(
		geometry.Pt(randInt(0, w), randInt(h, h+10)),
	))
	assert.Nil(t, canvas.Get(
		geometry.Pt(randInt(w, w+10), randInt(0, h)),
	))

	assert.Nil(t, canvas.Get(
		geometry.Pt(randInt(0, w), randInt(-10, 0)),
	))
	assert.Nil(t, canvas.Get(
		geometry.Pt(randInt(-10, 0), randInt(0, h)),
	))
}

func TestCanvas_Get_WithinBounds(t *testing.T) {
	w := rand.Int()
	h := rand.Int()
	canvas := MakeCanvas(geometry.Rect(0, 0, w, h))

	assert.NotNil(t, canvas.Get(
		geometry.Pt(randInt(0, w), randInt(0, h)),
	))
}

func TestCanvas_SubCanvas(t *testing.T) {
	w := rand.Int()
	h := rand.Int()
	bounds := geometry.Rect(0, 0, w, h)
	canvas := MakeCanvas(bounds)
	assert.Equal(t, bounds, canvas.Bounds)

	bounds = geometry.Rect(0, 0, w/2, h/2)
	subCanvas := canvas.SubCanvas(bounds)
	assert.Equal(t, bounds, subCanvas.Bounds)
	assert.Equal(t, subCanvas.cellGrid, canvas.cellGrid)
}
