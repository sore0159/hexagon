package hexagon

//  ========== HEX GRID ==========  //
/*
         S2		      S1

               0, 1
        -1, 1       1, 0
  S3           0, 0         S6
        -1, 0       1,-1
               0,-1
		S4			  S5
circumfremnce = radius*6 (1 for r = 0)
*/
var HEXDIRS = map[int]Coord{
	0: Coord{1, 0},
	1: Coord{0, 1},
	2: Coord{-1, 1},
	3: Coord{-1, 0},
	4: Coord{0, -1},
	5: Coord{1, -1},
}

// Sector returns which sector the coord lies in
// and has poorly thought out behavior if coord
// is on a straight line from the origin (use .Axis() before .Sector())
func (c Coord) Sector() int {
	if c[0] == 0 && c[1] == 0 {
		return 0
	}
	if c[0] >= 0 {
		if c[1] >= 0 {
			return 1
		} else if -1*c[1] > c[0] {
			return 5
		} else {
			return 6
		}
	} else if c[1] >= 0 {
		if -1*c[0] > c[1] {
			return 3
		} else {
			return 2
		}
	} else {
		return 4
	}
}

// Axis returns -1 if the coord is not on a straight line to
// the origin, otherwise returns the key for the HEXDIR from
// the origin to coord
func (c *Coord) Axis() int {
	if c[0] == 0 && c[0] == 0 {
		return -1
	}
	if (c[0] > 0 && c[1] > 0) || (c[0] < 0 && c[1] < 0) {
		return -1
	}
	if c[0] != 0 && c[1] != 0 {
		if c[0] != c[1]*-1 {
			return -1
		}
		if c[0] < 0 {
			return 2
		} else {
			return 5
		}
	}
	if c[0] == 0 {
		if c[1] < 0 {
			return 4
		} else {
			return 1
		}
	} else {
		if c[0] < 0 {
			return 3
		} else {
			return 0
		}
	}
}
