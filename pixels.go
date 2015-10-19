package hexagon

import (
	//	"fmt"
	"math"
)

var HEXSIZE = 10

type Pixel [2]float64

// Hex2Plane returns the x,y coords of the center of a hexagon
// chosen from a flat-side-up grid of "radius" sized hexagons,
// with the hexagon 0,0 having a center at 0,0
func Hex2Plane(radius int, coord [2]int) (plane [2]int) {
	x := 2 * radius * coord[0]
	// y^2 + (r/2)^2 = r^2
	// y = sqrt( r^2 - (r/2)^2 )
	y := float64(coord[1]*2-coord[0]) * math.Sqrt(float64(3*radius*radius)/4.0)
	return [2]int{x, int(y)}
}
