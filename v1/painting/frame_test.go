package painting

import (
	"image"
	"testing"
)

// ANCHOR TestNewFrame

type NewFrameOptions struct {
	p             Position
	width, height int
}

type NewFrameTest struct {
	input  NewFrameOptions
	output *Frame
}

var NewFrameTests []NewFrameTest = []NewFrameTest{
	NewFrameTest{
		input: NewFrameOptions{
			p:      image.Pt(0, 0),
			width:  0,
			height: 0,
		},
		output: &Frame{
			Position: Position{
				X: 0,
				Y: 0,
			},
			Patch: NewMatrix(0, 0),
		},
	},
	NewFrameTest{
		input: NewFrameOptions{
			p:      image.Pt(1, 0),
			width:  0,
			height: 0,
		},
		output: &Frame{
			Position: Position{
				X: 1,
				Y: 0,
			},
			Patch: NewMatrix(0, 0),
		},
	},
	NewFrameTest{
		input: NewFrameOptions{
			p:      image.Pt(0, 1),
			width:  0,
			height: 0,
		},
		output: &Frame{
			Position: Position{
				X: 0,
				Y: 1,
			},
			Patch: NewMatrix(0, 0),
		},
	},
	NewFrameTest{
		input: NewFrameOptions{
			p:      image.Pt(0, 0),
			width:  1,
			height: 0,
		},
		output: &Frame{
			Position: Position{
				X: 0,
				Y: 0,
			},
			Patch: NewMatrix(1, 0),
		},
	},
	NewFrameTest{
		input: NewFrameOptions{
			p:      image.Pt(0, 0),
			width:  0,
			height: 1,
		},
		output: &Frame{
			Position: Position{
				X: 0,
				Y: 0,
			},
			Patch: NewMatrix(0, 1),
		},
	},
	NewFrameTest{
		input: NewFrameOptions{
			p:      image.Pt(12, 5),
			width:  67,
			height: 890,
		},
		output: &Frame{
			Position: Position{
				X: 12,
				Y: 5,
			},
			Patch: NewMatrix(67, 890),
		},
	},
}

func TestNewFrame(t *testing.T) {
	for i := 0; i < len(NewFrameTests); i++ {
		test := NewFrameTests[i]
		frame := NewFrame(test.input.p, test.input.width, test.input.height)

		if !frame.isEqual(test.output) {
			t.Logf("Test n°%v: NewFrame doesn't return the expected result.", i)
			t.Logf("Getting: %+v", frame)
			t.Fatalf("Expecting: %+v", test.output)
		}
	}
}

// ANCHOR TestCanContain

type CanContainTest struct {
	input, other *Frame
	output       bool
}

var CanContainTests []CanContainTest = []CanContainTest{
	CanContainTest{
		input: NewFrame(
			image.Pt(0, 0),
			10, 10,
		),
		other: NewFrame(
			image.Pt(0, 0),
			10, 10,
		),
		output: true,
	},
	CanContainTest{
		input: NewFrame(
			image.Pt(0, 0),
			10, 10,
		),
		other: NewFrame(
			image.Pt(1, 1),
			10, 10,
		),
		output: false,
	},
	CanContainTest{
		input: NewFrame(
			image.Pt(0, 0),
			10, 10,
		),
		other: NewFrame(
			image.Pt(1, 1),
			5, 5,
		),
		output: true,
	},
	CanContainTest{
		input: NewFrame(
			image.Pt(0, 0),
			10, 10,
		),
		other: NewFrame(
			image.Pt(1, 0),
			10, 1,
		),
		output: false,
	},
	CanContainTest{
		input: NewFrame(
			image.Pt(0, 0),
			10, 10,
		),
		other: NewFrame(
			image.Pt(0, 1),
			1, 10,
		),
		output: false,
	},
	CanContainTest{
		input: NewFrame(
			image.Pt(0, 0),
			10, 10,
		),
		other: NewFrame(
			image.Pt(11, 1),
			0, 0,
		),
		output: false,
	},
	CanContainTest{
		input: NewFrame(
			image.Pt(0, 0),
			10, 10,
		),
		other: NewFrame(
			image.Pt(0, 11),
			0, 0,
		),
		output: false,
	},
}

func TestCanContain(t *testing.T) {
	for i := 0; i < len(CanContainTests); i++ {
		test := CanContainTests[i]
		frame := test.input
		other := test.other

		if can := frame.CanContain(other); can != test.output {
			t.Logf("Test n°%v: CanContain doesn't return the expected result.", i)
			t.Logf("Getting: %+v", can)
			t.Fatalf("Expecting: %+v", test.output)
		}
	}
}

// ANCHOR TestAdd

type AddTest struct {
	input, other *Frame
	output       *Frame
	err          bool
}

var AddTests []AddTest = []AddTest{
	AddTest{
		input: NewFrame(
			image.Pt(0, 0),
			2,
			3,
		),
		other: &Frame{
			Position: image.Pt(0, 0),
			Patch: &Matrix{
				M: [][]*Cell{
					[]*Cell{
						&Cell{
							Char: 'a',
						},
						&Cell{
							Char: 'b',
						},
					},
					[]*Cell{
						&CellDefault,
						&CellDefault,
					},
				},
			},
		},
		output: &Frame{
			Position: image.Pt(0, 0),
			Patch: &Matrix{
				M: [][]*Cell{
					[]*Cell{
						&Cell{
							Char: 'a',
						},
						&Cell{
							Char: 'b',
						},
					},
					[]*Cell{
						&CellDefault,
						&CellDefault,
					},
					[]*Cell{
						&CellDefault,
						&CellDefault,
					},
				},
			},
		},
	},
	AddTest{
		input: NewFrame(
			image.Pt(0, 0),
			2,
			4,
		),
		other: &Frame{
			Position: image.Pt(1, 2),
			Patch: &Matrix{
				M: [][]*Cell{
					[]*Cell{
						&Cell{
							Char: 'a',
						},
					},
					[]*Cell{
						&Cell{
							Char: 'b',
						},
					},
				},
			},
		},
		output: &Frame{
			Position: image.Pt(0, 0),
			Patch: &Matrix{
				M: [][]*Cell{
					[]*Cell{
						&CellDefault,
						&CellDefault,
					},
					[]*Cell{
						&CellDefault,
						&CellDefault,
					},
					[]*Cell{
						&CellDefault,
						&Cell{
							Char: 'a',
						},
					},
					[]*Cell{
						&CellDefault,
						&Cell{
							Char: 'b',
						},
					},
				},
			},
		},
	},
	AddTest{
		input: NewFrame(
			image.Pt(0, 0),
			0,
			0,
		),
		other: NewFrame(
			image.Pt(0, 0),
			0,
			0,
		),
		err: true,
	},
}

func TestAdd(t *testing.T) {
	for i := 0; i < len(AddTests); i++ {
		test := AddTests[i]
		frame := test.input
		other := test.other

		err := frame.Add(other)

		// error expected skip other checks
		if test.err && err != nil {
			continue
		}

		if err != nil {
			t.Logf("Test n°%v: Add return an error.", i)
			t.Logf("Getting: %+v", err)
			t.Fatalf("Expecting: %+v", nil)
		}

		if !frame.isEqual(test.output) {
			t.Logf("Test n°%v: Add doesn't return the expected result.", i)
			t.Logf("Getting: %+v", frame.Patch)
			t.Fatalf("Expecting: %+v", test.output.Patch)
		}
	}
}
