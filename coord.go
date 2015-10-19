package hexagon

import "fmt"

type Coord [2]int

func (c Coord) String() string {
	return fmt.Sprintf("HEX[x%d,y%d]", c[0], c[1])
}

func (c1 Coord) Add(c2 Coord) (c3 Coord) {
	return Coord{c1[0] + c2[0], c1[1] + c2[1]}
}

func (c1 Coord) Scale(n int) (c2 Coord) {
	return Coord{c1[0] * n, c1[1] * n}
}

func (c1 Coord) StepsTo(c2 Coord) int {
	x := c2[0] - c1[0]
	y := c2[1] - c1[1]
	if x == 0 {
		if y >= 0 {
			return y
		} else {
			return -1 * y
		}
	} else if y == 0 {
		if x >= 0 {
			return x
		} else {
			return -1 * x
		}
	}
	if x < 0 && y < 0 {
		if x < y {
			return -1 * x
		} else {
			return -1 * y
		}
	} else if x > 0 && y > 0 {
		if x > y {
			return x
		} else {
			return y
		}
	}
	if x < 0 {
		x *= -1
	}
	if y < 0 {
		y *= -1
	}
	return y + x
}

func (c Coord) Polar() Polar {
	if c[0] == 0 && c[1] == 0 {
		return Polar{0, 0}
	}
	r := c.StepsTo(Coord{0, 0})
	circ := 6 * r
	for theta := 0; theta < circ; theta++ {
		test := Polar{r, theta}
		if test.Coord() == c {
			return test
		}
	}
	return Polar{-1, -1}
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
		leg = leg.Scale(r)
		legDir := HEXDIRS[(i+2)%6]
		for j := 0; j < r; j++ {
			path = append(path, leg.Add(legDir.Scale(j)))
		}
	}
	return path
}
