package hexagon

import "fmt"

type Coord [2]int

func (c Coord) String() string {
	return fmt.Sprintf("HEX[x%d,y%d]", c[0], c[1])
}

func (c1 Coord) Add(c2 Coord) (c3 Coord) {
	return Coord{c1[0] + c2[0], c1[1] + c2[1]}
}

func (c1 Coord) Subtract(c2 Coord) (c3 Coord) {
	return Coord{c1[0] - c2[0], c1[1] - c2[1]}
}

func (c1 Coord) Scale(n int) (c2 Coord) {
	return Coord{c1[0] * n, c1[1] * n}
}

func (c1 Coord) StepsTo(c2 Coord) int {
	c1z := -c1[0] - c1[1]
	c2z := -c2[0] - c2[1]
	dx := Abs(c2[0] - c1[0])
	dy := Abs(c2[1] - c1[1])
	dz := Abs(c2z - c1z)
	return (dx + dy + dz) / 2
}

func (c Coord) Polar() Polar {
	if c[0] == 0 && c[1] == 0 {
		return Polar{0, 0}
	}
	r := c.StepsTo(Coord{0, 0})
	if leg := c.Axis(); leg != -1 {
		return Polar{r, r * leg}
	}
	sector := c.Sector()
	legD := HEXDIRS[sector-1].Scale(r)
	extD := HEXDIRS[(sector+1)%6]
	// legD[0]+X*extD[0] == c[0]
	// legD[1]+X*extD[1] == c[1]
	// X*(extD[0]-extD[1]) == c[0]-c[1] - (legD[0] -legD[1])
	theta := r*(sector-1) + (c[0]-c[1]-(legD[0]-legD[1]))/(extD[0]-extD[1])
	return Polar{r, theta}
}

func (c Coord) Ring(r int) []Coord {
	if r < 0 {
		return nil
	} else if r == 0 {
		return []Coord{c}
	}
	path := []Coord{}
	for i := 0; i < 6; i++ {
		leg := HEXDIRS[i]
		leg = c.Add(leg.Scale(r))
		legDir := HEXDIRS[(i+2)%6]
		for j := 0; j < r; j++ {
			path = append(path, leg.Add(legDir.Scale(j)))
		}
	}
	return path
}
