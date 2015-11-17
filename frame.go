package hexagon

import "log"

// SetFrame sets the boundaries for VisList()
// If the given values are not geometrically
// sound it flip them as needed.
// Note that the frame does _not_ need to include
// the CenterPix: CenterPix is only used for
// hex/pixel translation.
func (v *Viewport) SetFrame(x1, y1, x2, y2 float64) {
	log.Println("V:", v.ULCorner, v.LRCorner, x1, y1, x2, y2)
	if x1 < x2 {
		v.ULCorner[0] = x1
		v.LRCorner[0] = x2
	} else {
		v.ULCorner[0] = x2
		v.LRCorner[0] = x1
	}
	if y1 < y2 {
		v.ULCorner[1] = y1
		v.LRCorner[1] = y2
	} else {
		v.ULCorner[1] = y2
		v.LRCorner[1] = y1
	}
	log.Println("V2:", v.ULCorner, v.LRCorner)
}

// VisList returns a list of all Coords within
// the boundaries of the Frame as set by SetFrame
// VisList errs on the side of inclusion: you should
// always get a hex in the frame, and sometimes
// (but only rarely) get a hex outside the frame
// VisList includes hexagons partially visible, and
// has a one-or-two hexgon inefficiency sometimes
//       (extra rows included sometimes
//         have an unneeded hexagon)
// but generally does a good job of including only
// those hexes it needs to
func (v *Viewport) VisList() (surface []Coord) {
	if v.Flattop {
		return v.flatVisList()
	} else {
		return v.flipVisList()
	}
}

func (v *Viewport) flatVisList() (surface []Coord) {
	ulHex := v.HexContaining(v.ULCorner)
	urHex := v.HexContaining(Pixel{v.LRCorner[0], v.ULCorner[1]})
	firstCol, lastCol := ulHex[0], urHex[0]
	if v.CenterOf(ulHex)[0]-v.ULCorner[0] > v.HexW*.25 {
		firstCol--
	}
	if v.LRCorner[0]-v.CenterOf(urHex)[0] > v.HexW*.25 {
		lastCol++
	}
	for col := firstCol; col <= lastCol; col++ {
		surface = append(surface, v.flatVisColumn(col)...)
	}
	return
}

func (v *Viewport) flipVisList() (surface []Coord) {
	ulHex := v.HexContaining(v.ULCorner)
	llHex := v.HexContaining(Pixel{v.ULCorner[0], v.LRCorner[1]})
	firstRow, lastRow := ulHex[1], llHex[1]
	if v.CenterOf(ulHex)[1]-v.ULCorner[1] > v.HexW*.25 {
		firstRow++
	}
	if v.LRCorner[1]-v.CenterOf(llHex)[1] > v.HexW*.25 {
		lastRow--
	}
	for row := firstRow; row >= lastRow; row-- {
		surface = append(surface, v.flipVisRow(row)...)
	}
	return
}

func (v *Viewport) flipVisRow(y int) (row []Coord) {
	row = make([]Coord, 0)

	pxY := v.flipCenterOf(Coord{0, y})[1]
	start := v.HexContaining(Pixel{v.ULCorner[0], pxY})
	stop := v.HexContaining(Pixel{v.LRCorner[0], pxY})
	for x := start[0]; x <= stop[0]; x++ {
		row = append(row, Coord{x, start[1]})
	}
	return row
}

func (v *Viewport) flatVisColumn(x int) (col []Coord) {
	col = make([]Coord, 0)
	pxX := v.flatCenterOf(Coord{x, 0})[0]
	start := v.HexContaining(Pixel{pxX, v.ULCorner[1]})
	stop := v.HexContaining(Pixel{pxX, v.LRCorner[1]})
	for y := start[1]; y >= stop[1]; y-- {
		col = append(col, Coord{start[0], y})
	}
	return col
}
