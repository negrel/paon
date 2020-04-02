package style

import "log"

// Grid represent a grid of terminal screen cell.
type Grid []Row

// AppendRow method append the given row.
func (g *Grid) appendRow(row Row) {
	*g = append(*g, row)
}

// AppendAt method append the given row extension
// to the n row. If the row doesn't exist, panic.
func (g *Grid) appendAt(ext Row, n int) {
	if 0 > n || n >= len(*g) {
		log.Panic("Impossible to append to a non-existing row.")
	}

	(*g)[n] = append((*g)[n], ext...)
}
