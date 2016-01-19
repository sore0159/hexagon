package hexagon

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -1 * x
}

func Dir(x int) int {
	if x < 0 {
		return -1
	} else {
		return 1
	}
}

func HexArea(rad int) int {
	if rad < 0 {
		return 0
	} else if rad == 0 {
		return 1
	}
	return 1 + 3*(rad*(rad+1))
}
