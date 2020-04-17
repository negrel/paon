// package render

// import (
// 	"testing"

// 	"github.com/negrel/ginger/v2/style"
// )

// type size struct {
// 	width, height int
// }

// // ANCHOR TestNewMatrix

// type NewMatrixTest struct {
// 	input size

// 	output *Matrix
// }

// var NewMatrixTests []NewMatrixTest = []NewMatrixTest{
// 	{
// 		input: size{
// 			width:  0,
// 			height: 0,
// 		},
// 		output: &Matrix{
// 			M: [][]*Cell{},
// 		},
// 	},

// 	{
// 		input: size{
// 			width:  1,
// 			height: 2,
// 		},
// 		output: &Matrix{
// 			M: [][]*Cell{
// 				{
// 					&CellDefault,
// 				},
// 				{
// 					&CellDefault,
// 				},
// 			},
// 		},
// 	},
// }

// func TestNewMatrix(t *testing.T) {
// 	for i := 0; i < len(NewMatrixTests); i++ {
// 		test := NewMatrixTests[i]
// 		matrix := NewMatrix(test.input.width, test.input.height)

// 		if !matrix.isEqual(test.output) {
// 			t.Logf("Test n°%v: NewMatrix doesn't return the expected result.", i)
// 			t.Logf("Getting: %+v", matrix)
// 			t.Fatalf("Expecting: %+v", test.output)
// 		}
// 	}
// }

// type WidthAndHeightTest struct {
// 	input  *Matrix
// 	output size
// }

// var WidthAndHeightTests []WidthAndHeightTest = []WidthAndHeightTest{
// 	{
// 		input: NewMatrix(0, 0),
// 		output: size{
// 			width:  0,
// 			height: 0,
// 		},
// 	},
// 	// NOTE Width can't be equal to one if height is 0
// 	{
// 		input: NewMatrix(1, 0),
// 		output: size{
// 			width:  0,
// 			height: 0,
// 		},
// 	},
// 	{
// 		input: NewMatrix(0, 1),
// 		output: size{
// 			width:  0,
// 			height: 1,
// 		},
// 	},
// 	{
// 		input: NewMatrix(7894, 2136),
// 		output: size{
// 			width:  7894,
// 			height: 2136,
// 		},
// 	},
// }

// func TestWidthAndHeight(t *testing.T) {
// 	for i := 0; i < len(WidthAndHeightTests); i++ {
// 		test := WidthAndHeightTests[i]
// 		matrix := test.input

// 		if mWidth := matrix.Width(); mWidth != test.output.width {
// 			t.Logf("Test n°%v: Width doesn't return the expected result.", i)
// 			t.Logf("Getting: %+v", mWidth)
// 			t.Fatalf("Expecting: %+v", test.output.width)
// 		}

// 		if mHeight := matrix.Height(); mHeight != test.output.height {
// 			t.Logf("Test n°%v: Height doesn't return the expected result.", i)
// 			t.Logf("Getting: %+v", mHeight)
// 			t.Fatalf("Expecting: %+v", test.output.height)
// 		}
// 	}
// }

// // ANCHOR TestIsEmpty

// type IsEmptyTest struct {
// 	input *Matrix

// 	output bool
// }

// var IsEmptyTests []IsEmptyTest = []IsEmptyTest{
// 	{
// 		input:  NewMatrix(0, 0),
// 		output: true,
// 	},
// 	{
// 		input:  NewMatrix(7657, 1000),
// 		output: false,
// 	},
// }

// func TestIsEmpty(t *testing.T) {
// 	for i := 0; i < len(IsEmptyTests); i++ {
// 		test := IsEmptyTests[i]

// 		if isEmpty := test.input.isEmpty(); isEmpty != test.output {
// 			t.Logf("Test n°%v: isEmpty doesn't return the expected result.", i)
// 			t.Logf("Getting: %+v", isEmpty)
// 			t.Fatalf("Expecting: %+v", test.output)
// 		}
// 	}
// }

// // ANCHOR TestIsEqual

// type IsEqualTest struct {
// 	input, other *Matrix

// 	output bool
// }

// var IsEqualTests []IsEqualTest = []IsEqualTest{
// 	{
// 		input:  NewMatrix(0, 0),
// 		other:  NewMatrix(0, 0),
// 		output: true,
// 	},
// 	{
// 		input:  NewMatrix(7657, 1000),
// 		other:  NewMatrix(7657, 1000),
// 		output: true,
// 	},
// 	{
// 		input:  NewMatrix(7657, 1000),
// 		other:  NewMatrix(7656, 1001),
// 		output: false,
// 	},
// 	{
// 		input: &Matrix{
// 			M: [][]*Cell{
// 				{
// 					&CellDefault,
// 				},
// 				{
// 					&CellDefault,
// 				},
// 			},
// 		},
// 		other: &Matrix{
// 			M: [][]*Cell{
// 				{
// 					&Cell{
// 						Theme: style.ThemeDefault(),
// 						Char:  0,
// 					},
// 				},
// 				{
// 					&Cell{
// 						Theme: style.ThemeDefault(),
// 						Char:  0,
// 					},
// 				},
// 			},
// 		},
// 		output: true,
// 	},
// 	{
// 		input: &Matrix{
// 			M: [][]*Cell{
// 				{
// 					&CellDefault,
// 				},
// 				{
// 					&CellDefault,
// 				},
// 			},
// 		},
// 		other: &Matrix{
// 			M: [][]*Cell{
// 				{
// 					&Cell{
// 						Theme: &style.Theme{
// 							Foreground: 0x000000,
// 							Background: style.ColorDefault,
// 						},
// 						Char: '0',
// 					},
// 				},
// 				{
// 					&Cell{
// 						Theme: &style.Theme{
// 							Foreground: style.ColorDefault,
// 							Background: 0x000000,
// 						},
// 						Char: '0',
// 					},
// 				},
// 			},
// 		},
// 		output: false,
// 	},
// 	{
// 		input: &Matrix{
// 			M: [][]*Cell{
// 				{
// 					&Cell{
// 						Char: 'a',
// 					},
// 				},
// 				{
// 					&Cell{
// 						Char: 'b',
// 					},
// 				},
// 			},
// 		},
// 		other: &Matrix{
// 			M: [][]*Cell{
// 				{
// 					&Cell{
// 						Char: 'a',
// 					},
// 				},
// 				{
// 					&Cell{
// 						Char: 'b',
// 					},
// 				},
// 			},
// 		},
// 		output: true,
// 	},
// 	{
// 		input: &Matrix{
// 			M: [][]*Cell{
// 				{
// 					&Cell{
// 						Char: 'a',
// 					},
// 				},
// 				{
// 					&Cell{
// 						Char: 'b',
// 					},
// 				},
// 			},
// 		},
// 		other: &Matrix{
// 			M: [][]*Cell{
// 				{
// 					&Cell{
// 						Char: 'c',
// 					},
// 				},
// 				{
// 					&Cell{
// 						Char: 'd',
// 					},
// 				},
// 			},
// 		},
// 		output: false,
// 	},
// }

// func TestIsEqual(t *testing.T) {
// 	for i := 0; i < len(IsEqualTests); i++ {
// 		test := IsEqualTests[i]

// 		if isEqual := test.input.isEqual(test.other); isEqual != test.output {
// 			t.Logf("Test n°%v: isEqual doesn't return the expected result.", i)
// 			t.Logf("Getting: %+v", isEqual)
// 			t.Fatalf("Expecting: %+v", test.output)
// 		}
// 	}
// }
