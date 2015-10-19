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
