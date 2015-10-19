package hexagon

// HexPath creates a path of 1 step grid coords, starting with A and ending with B
// HexPath uses HexPathSteps to calculate the least steps to take, and
// then uses StepSplit to interweave the steps as evenly as possible

func (c1 Coord) PathTo(c2 Coord) []Coord {
	d := c1.StepsTo(c2)
	if d == 0 {
		return []Coord{c1}
	} else if d == 1 {
		return []Coord{c1, c2}
	}
	steps := PathSteps(c1, c2) // num steps HEXDIR 0, 1, 2
	var numLarge, numSmall int
	var sModX, sModY, lModX, lModY int
	if Abs(steps[0]) >= Abs(steps[1]) && Abs(steps[0]) >= Abs(steps[2]) {
		numLarge = Abs(steps[0])
		lModX, lModY = Dir(steps[0]), 0
		if steps[1] == 0 {
			numSmall = Abs(steps[2])
			sModX, sModY = Dir(steps[2]), Dir(steps[2])
		} else {
			numSmall = Abs(steps[1])
			sModX, sModY = 0, Dir(steps[1])
		}
	} else if Abs(steps[1]) >= Abs(steps[2]) {
		numLarge = Abs(steps[1])
		lModX, lModY = 0, Dir(steps[1])
		if steps[0] == 0 {
			numSmall = Abs(steps[2])
			sModX, sModY = Dir(steps[2]), Dir(steps[2])
		} else {
			numSmall = Abs(steps[0])
			sModX, sModY = Dir(steps[0]), 0
		}
	} else {
		numLarge = Abs(steps[2])
		lModX, lModY = Dir(steps[2]), Dir(steps[2])
		if steps[0] == 0 {
			numSmall = Abs(steps[1])
			sModX, sModY = 0, Dir(steps[1])
		} else {
			numSmall = Abs(steps[0])
			sModX, sModY = Dir(steps[0]), 0
		}
	}
	useSmall := StepSplit(numLarge, numSmall)
	path := []Coord{c1}
	stepper := c1
	for count := 0; count < d; count++ {
		var modX, modY int
		if useSmall[count] {
			modX, modY = sModX, sModY
		} else {
			modX, modY = lModX, lModY
		}
		stepper = stepper.Add(Coord{modX, modY})
		path = append(path, stepper)
	}
	return path
}

// PathSteps calculates the minimum steps from a to b
// PathSteps returns a [3]int with the count of:
// StepsRight(0,1) StepsUp(1,0) StepsUpRight(1,1)
// with negative numbers for backward steps
func PathSteps(a, b Coord) (steps [3]int) {
	//steps right up upRight
	x := b[0] - a[0]
	y := b[1] - a[1]
	if x == 0 {
		steps[1] = y
	} else if y == 0 {
		steps[0] = x
	} else if x < 0 && y < 0 {
		if x < y {
			steps[2] = y
			steps[0] = x - y
		} else {
			steps[2] = x
			steps[1] = y - x
		}
	} else if x > 0 && y > 0 {
		if x > y {
			steps[2] = y
			steps[0] = x - y
		} else {
			steps[2] = x
			steps[1] = y - x
		}
	} else {
		steps[0] = x
		steps[1] = y
	}
	return
}

// StepSplit figures out how best to weave two types of steps
// It takes two ints representing counts of each steptype and is
// indifferent to what those steptypes are
// It returns a bool slice to use for determining which type to
// use in a sequence
func StepSplit(larger, smaller int) (useSmall []bool) {
	//fmt.Println("Got L,S:", larger, smaller)
	useSmall = make([]bool, larger+smaller)
	if smaller <= 0 {
		return
	}
	if larger == smaller {
		for i, _ := range useSmall {
			useSmall[i] = i%2 == 1
		}
		return
	}
	beat := (larger + smaller) / (smaller + 1)
	beats := make([]int, smaller)
	for i, _ := range beats {
		beats[i] = beat*(i+1) - 1
	}
	left := (larger + smaller) - beat*smaller
	// left-beat = larger+smaller - beat*(smaller+1)
	// l-b = sum - (< sum)
	// left > beat
	midbeat := len(beats) / 2
	//fmt.Println("Beat:", beat, "beats:", beats, "midbeat:", midbeat)
	for i := 0; left > (beat - 1); left-- {
		boost := midbeat + i
		if i <= 0 {
			i *= -1
			i++
		} else {
			i *= -1
		}
		for j, val := range beats {
			if j >= boost {
				beats[j] = val + 1
			}
		}
	}
	//fmt.Println("beats:", beats)
	for i, _ := range useSmall {
		if len(beats) > 0 && i == beats[0] {
			useSmall[i] = true
			if len(beats) > 1 {
				beats = beats[1:]
			} else {
				beats = nil
			}
		}
	}
	//fmt.Println("Giving:", useSmall)
	/*var sm, lr int
	for _, val := range useSmall {
		if val {
			sm++
		} else {
			lr++
		}
	}
	fmt.Println("Sums:L,S:", lr, sm)
	*/
	/*path := []int{0}
	dir := 1
	for i, val := range useSmall {
		path[len(path)-1] += dir
		if i < len(useSmall)-1 {
			if useSmall[i+1] != val {
				path = append(path, 0)
				dir *= -1
			}
		}
	}
	fmt.Println("Switches:", path)
	*/
	return
}
