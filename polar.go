package hexagon

import "fmt"

type Polar [2]int // radius, theta

func (p Polar) String() string {
	return fmt.Sprintf("POLAR[r%d,t%d]", p[0], p[1])
}

func (p Polar) Coord() Coord {
	if p[0] == 0 {
		return Coord{0, 0}
	}
	radius, theta := p[0], p[1]
	circ := radius * 6
	for theta > circ {
		theta -= circ
	}
	for theta < 0 {
		theta += circ
	}
	// Which leg of the hexagon are we on?
	// (legs have length = radius)
	leg := theta / radius
	// How far along that leg are we?
	extra := theta % radius
	// Go out along the axis to the start of your leg
	extend := HEXDIRS[leg]
	extend = [2]int{extend[0] * radius, extend[1] * radius}
	extend = extend.Scale(radius)
	// Go along the leg the appropriate distance
	legDir := HEXDIRS[(leg+2)%6]
	legDir = legDir.Scale(extra)
	// Viola!
	return extend.Add(legDir)
}
