package hexagon

import (
	"fmt"
	"math"
)

/*
            ____
           /\  /\
	  ____/__\/__\        |\
	 /\  /\  /\  /    r/2 | \  r
	/__\/__\/__\/     *3^ |  \
    \  /\  /          1/2 |   \
     \/__\/               |____\
                            r/2
	h = r*sqrt(3)
	w = 2r
*/

type Pixel [2]float64

func (p Pixel) String() string {
	return fmt.Sprintf("PIXEL[x%.2f,y%.2f]", p[0], p[1])
}

func (p Pixel) Add(p2 Pixel) (p3 Pixel) {
	return Pixel{p[0] + p2[0], p[1] + p2[1]}
}

func (p Pixel) Subtract(p2 Pixel) (p3 Pixel) {
	return Pixel{p[0] - p2[0], p[1] - p2[1]}
}

func (p Pixel) Scale(s float64) (p2 Pixel) {
	return Pixel{p[0] * s, p[1] * s}
}

func (p Pixel) Flip() Pixel {
	return Pixel{p[1], p[0]}
}

// Viewport is a tool for displaying a hexagon grid
// CenterPix is the pixel at the center of CenterHex
// The y-axis is inverted for use in images
type Viewport struct {
	HexR      float64
	HexH      float64
	HexW      float64
	CenterHex Coord
	CenterPix Pixel
	Width     float64
	Height    float64
}

func MakeViewport(r float64) *Viewport {
	return &Viewport{
		HexR: r,
		HexH: r * math.Sqrt(3),
		HexW: r * 2,
	}
}

func (v *Viewport) SetRadius(r float64) {
	v.HexR = r
	v.HexH = r * math.Sqrt(3)
	v.HexW = r * 2
}

// CenterOf returns the pixel at the center of the hexagon c
func (v *Viewport) CenterOf(c Coord) Pixel {
	shift := c.Subtract(v.CenterHex)
	//radius := v.HexR
	x := .75 * v.HexW * float64(shift[0])
	//x := 1.5 * radius * float64(shift[0])
	// y^2 + (r/2)^2 = r^2
	// y = sqrt( r^2 - (r/2)^2 ) =  sqr(3)*r/2
	y := float64(shift[1]*2+shift[0]) * .5 * (v.HexH)
	return v.CenterPix.Add(Pixel{x, -y})
}

// Corners returns the pixels at the corners of the hexagon at c
// starting at 3 o'clock and rotating counter-clockwise
func (v *Viewport) CornersOf(c Coord) (corners [6]Pixel) {
	center := v.CenterOf(c)
	//radius := v.HexR
	dx := v.HexW * .5
	//corners[0] = center.Add(Pixel{radius, 0})
	//corners[3] = center.Add(Pixel{-1 * radius, 0})
	corners[0] = center.Add(Pixel{dx, 0})
	corners[3] = center.Add(Pixel{-dx, 0})
	// r^2 = (r/2)^2 + y^2
	// y^2 = 3 r^2 / 4
	//x := .5 * radius
	x := .5 * dx
	y := .5 * v.HexH
	corners[1] = center.Add(Pixel{x, -y})
	corners[5] = center.Add(Pixel{x, y})
	corners[2] = center.Add(Pixel{-x, -y})
	corners[4] = center.Add(Pixel{-x, y})
	return
}

// HexContaining returns the coord of the hexagon containing
// the given pixel
// There is border ambugiuity of one pixel, handled by rounding
func (v *Viewport) HexContaining(p Pixel) (c Coord) {
	//		 __
	//		| \|  width =  1.5*r
	//box = |_/|  height = r*sqrt(3)
	// upper left corner:  centerPix[0] - .5r, centerPix[1] + h/2
	r := v.HexW * .5
	h := v.HexH
	w := 1.5 * r
	// Which box?
	clickX := p[0] - (v.CenterPix[0] - .5*r)
	x := int(math.Floor(clickX / w))
	clickY := p[1] - (v.CenterPix[1] - .5*h) + (.5 * h * float64(x))
	y := -int(math.Floor(clickY / h))
	c = v.CenterHex.Add(Coord{x, y})
	// Where in the box?
	center := v.CenterOf(c)
	boxX := p[0] - center[0]
	boxY := p[1] - center[1]
	if boxX <= r/2 {
		return
	}
	boxX -= r / 2
	if boxY >= 0 {
		if r/2-boxX < r*boxY/h {
			c = c.Add(Coord{1, -1})
		}
	} else {
		if r/2-boxX < -r*boxY/h {
			c = c.Add(Coord{1, 0})
		}
	}
	return
}
