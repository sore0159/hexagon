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
// The y-axis is inverted for use in images (hex 0,1
// has lower valued y pixels than hex 0,0)
type Viewport struct {
	HexR      float64
	HexH      float64
	HexW      float64
	CenterHex Coord
	CenterPix Pixel
	Flattop   bool
	Isometric bool
	ULCorner  Pixel
	LRCorner  Pixel
}

func MakeViewport(r float64, flattop bool, isometric bool) *Viewport {
	h := r * math.Sqrt(3)
	w := 2 * r
	if isometric {
		if flattop {
			h *= .75
		} else {
			w *= .75
		}
	}
	return &Viewport{
		HexR:      r,
		HexH:      h,
		HexW:      w,
		Flattop:   flattop,
		Isometric: isometric,
	}
}

func (v *Viewport) Anchor(c Coord, p Pixel) {
	v.CenterHex = c
	v.CenterPix = p
}

func (v *Viewport) SetFrame(upperLeft, lowerRight Pixel) {
	v.ULCorner = upperLeft
	v.LRCorner = lowerRight
}

func (v *Viewport) SetRadius(r float64) {
	v.HexR = r
	h := r * math.Sqrt(3)
	w := r * 2
	if v.Isometric {
		if v.Flattop {
			h *= .75
		} else {
			w *= .75
		}
	}
	v.HexH = h
	v.HexW = w
}

func (v *Viewport) SetFlattop(setTo bool) {
	if v.Flattop == setTo {
		return
	}
	v.Flattop = setTo
	if v.Isometric {
		if setTo {
			v.HexH *= .75
			v.HexW /= .75
		} else {
			v.HexH /= .75
			v.HexW *= .75
		}
	}
}

func (v *Viewport) SetIso(setTo bool) {
	if v.Isometric == setTo {
		return
	}
	if setTo {
		if v.Flattop {
			v.HexH *= .75
		} else {
			v.HexW *= .75
		}
	} else {
		if v.Flattop {
			v.HexH /= .75
		} else {
			v.HexW /= .75
		}
	}
	v.Isometric = setTo
}

// CenterOf returns the pixel at the center of the hexagon c
func (v *Viewport) CenterOf(c Coord) Pixel {
	if v.Flattop {
		return v.flatCenterOf(c)
	}
	return v.flipCenterOf(c)
}

func (v *Viewport) flatCenterOf(c Coord) Pixel {
	shift := c.Subtract(v.CenterHex)
	//radius := v.HexR
	x := .75 * v.HexW * float64(shift[0])
	//x := 1.5 * radius * float64(shift[0])
	// y^2 + (r/2)^2 = r^2
	// y = sqrt( r^2 - (r/2)^2 ) =  sqr(3)*r/2
	y := float64(shift[1]*2+shift[0]) * .5 * (v.HexH)
	return v.CenterPix.Add(Pixel{x, -y})
}

func (v *Viewport) flipCenterOf(c Coord) Pixel {
	shift := c.Subtract(v.CenterHex)
	//radius := v.HexR
	y := .75 * v.HexW * float64(shift[1])
	//x := 1.5 * radius * float64(shift[0])
	// y^2 + (r/2)^2 = r^2
	// y = sqrt( r^2 - (r/2)^2 ) =  sqr(3)*r/2
	x := float64(shift[0]*2+shift[1]) * .5 * (v.HexH)
	return v.CenterPix.Add(Pixel{x, -y})
}

// Corners returns the pixels at the corners of the hexagon at c
// starting at 3 o'clock and rotating counter-clockwise
func (v *Viewport) CornersOf(c Coord) [6]Pixel {
	if v.Flattop {
		return v.flatCornersOf(c)
	}
	return v.flipCornersOf(c)
}

func (v *Viewport) flatCornersOf(c Coord) (corners [6]Pixel) {
	center := v.flatCenterOf(c)
	//radius := v.HexR
	dx := v.HexW * .5
	//corners[0] = center.Add(Pixel{radius, 0})
	//corners[3] = center.Add(Pixel{-1 * radius, 0})
	// r^2 = (r/2)^2 + y^2
	// y^2 = 3 r^2 / 4
	//x := .5 * radius
	x := .5 * dx
	y := .5 * v.HexH
	corners[0] = center.Add(Pixel{dx, 0})
	corners[3] = center.Add(Pixel{-dx, 0})
	corners[1] = center.Add(Pixel{x, -y})
	corners[5] = center.Add(Pixel{x, y})
	corners[2] = center.Add(Pixel{-x, -y})
	corners[4] = center.Add(Pixel{-x, y})
	return
}

func (v *Viewport) flipCornersOf(c Coord) (corners [6]Pixel) {
	center := v.flipCenterOf(c)
	//radius := v.HexR
	dy := v.HexW * .5
	//corners[0] = center.Add(Pixel{radius, 0})
	//corners[3] = center.Add(Pixel{-1 * radius, 0})
	// r^2 = (r/2)^2 + y^2
	// y^2 = 3 r^2 / 4
	//x := .5 * radius
	y := .5 * dy
	x := .5 * v.HexH
	corners[0] = center.Add(Pixel{0, -dy})
	corners[1] = center.Add(Pixel{x, -y})
	corners[2] = center.Add(Pixel{x, y})
	corners[3] = center.Add(Pixel{0, dy})
	corners[4] = center.Add(Pixel{-x, y})
	corners[5] = center.Add(Pixel{-x, -y})
	return
}

// HexContaining returns the coord of the hexagon containing
// the given pixel
// There is border ambugiuity of one pixel, handled by rounding
func (v *Viewport) HexContaining(p Pixel) Coord {
	if v.Flattop {
		return v.flatHexContaining(p)
	}
	return v.flipHexContaining(p)
}

func (v *Viewport) flatHexContaining(p Pixel) (c Coord) {
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
	center := v.flatCenterOf(c)
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

func (v *Viewport) flipHexContaining(p Pixel) (c Coord) {
	//   __
	//  |  |  boxH = hexW*.75
	//   \/   boxW = hexH
	vertR := v.HexW * .5
	h := .75 * v.HexW
	w := v.HexH
	// Which box?
	// CenterBox UL = center[0]- .5*w,  center[1]- .5*vertR
	yDist := (p[1] - (v.CenterPix[1] - .5*vertR)) / h
	y := -int(math.Floor(yDist))
	xDist := (p[0] - (v.CenterPix[0] - .5*w) - .5*w*float64(y)) / w
	x := int(math.Floor(xDist))
	c = v.CenterHex.Add(Coord{x, y})
	center := v.flipCenterOf(c)
	boxX := p[0] - center[0]
	boxY := p[1] - center[1]
	if boxY <= vertR/2 {
		return
	}
	boxY -= vertR / 2
	r := v.HexR
	_ = r
	if boxX >= 0 {
		if vertR/2-boxY < vertR*boxX/w {
			c = c.Add(Coord{1, -1})
		}
	} else {
		if vertR/2-boxY < -vertR*boxX/w {
			c = c.Add(Coord{1, 0})
		}
	}
	return
}
