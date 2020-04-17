package render

// Matrix is a matrix of terminal cells. It is use
// for painting frame on the screen.
type Matrix struct {
	M [][]*Cell
}

// NewMatrix return a new empty matrix with the given dimension.
func NewMatrix(width, height int) *Matrix {
	matrix := make([][]*Cell, height, height)

	for i := 0; i < height; i++ {
		row := make([]*Cell, width, width)

		for j := 0; j < width; j++ {
			cell := &CellDefault

			row[j] = cell
		}

		matrix[i] = row
	}

	return &Matrix{
		M: matrix,
	}
}

/*****************************************************
 ***************** GETTERS & SETTERS *****************
 *****************************************************/
// ANCHOR Getters & setter

// Height return the matrix height
func (m *Matrix) Height() int {
	return len(m.M)
}

// Width return the matrix width.
func (m *Matrix) Width() int {
	if m.Height() > 0 {
		return len(m.M[0])
	}

	return 0

}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

func (m *Matrix) isEmpty() bool {
	if m.Height() == 0 || m.Width() == 0 {
		return true
	}

	return false
}

func (m *Matrix) isEqual(o *Matrix) bool {
	mWidth, mHeight := m.Width(), m.Height()
	oWidth, oHeight := o.Width(), o.Height()

	if mWidth != oWidth ||
		mHeight != oHeight {
		return false
	}

	for i := 0; i < mHeight; i++ {
		for j := 0; j < mWidth; j++ {
			if !m.M[i][j].isEqual(o.M[i][j]) {
				return false
			}
		}
	}

	return true
}

// isValid check that each row have the same length.
func (m *Matrix) isValid() bool {
	width := m.Width()
	height := m.Height()

	if height > 0 && width > 0 {
		for i := 0; i < height; i++ {
			row := m.M[i]
			for j := 0; j < width; j++ {
				if len(row) != width {
					return false
				}
			}
		}
	}

	return true
}
