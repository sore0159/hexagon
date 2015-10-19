package hexagon

//  ========== HEX GRID ==========  //
/*

       0, 1
 -1, 0       1, 1
       0, 0
 -1,-1       1, 0
       0,-1

circumfremnce = radius*6 (1 for r = 0)
area =
*/
var HEXDIRS = map[int]Coord{
	0: Coord{1, 0},
	1: Coord{1, 1},
	2: Coord{0, 1},
	3: Coord{-1, 0},
	4: Coord{-1, -1},
	5: Coord{0, -1},
}
