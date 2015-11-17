package hexagon

import (
	"fmt"
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
