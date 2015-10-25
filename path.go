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
	stepsN, stepD := PathSteps(c1, c2) // num steps HEXDIR 0, 1, 2
	path := []Coord{c1}
	if len(stepsN) == 1 {
		for i := 1; i <= d; i++ {
			path = append(path, c1.Add(stepD[0].Scale(i)))
		}
		return path
	}
	var useSmall []bool
	var firstLarge bool
	if stepsN[0] > stepsN[1] {
		useSmall = StepSplit(stepsN[0], stepsN[1])
		firstLarge = true
	} else {
		useSmall = StepSplit(stepsN[1], stepsN[0])
	}
	stepper := c1
	for count := 0; count < d; count++ {
		if (firstLarge && useSmall[count]) || (!firstLarge && !useSmall[count]) {
			stepper = stepper.Add(stepD[1])
		} else {
			stepper = stepper.Add(stepD[0])
		}
		path = append(path, stepper)
	}
	return path
}

// PathSteps calculates the minimum steps from a to b
func PathSteps(a, b Coord) (num []int, dir []Coord) {
	if a == b {
		return nil, nil
	}
	adjB := b.Subtract(a)
	if axis := adjB.Axis(); axis != -1 {
		d := a.StepsTo(b)
		return []int{d}, []Coord{HEXDIRS[axis]}
	}
	sector := adjB.Sector()
	dir = []Coord{HEXDIRS[sector-1], HEXDIRS[sector%6]}
	switch sector {
	/*0: Coord{1, 0},
	1: Coord{0, 1},
	2: Coord{-1, 1},
	3: Coord{-1, 0},
	4: Coord{0, -1},
	5: Coord{1, -1},*/
	case 1:
		return []int{adjB[0], adjB[1]}, dir
	case 2:
		return []int{adjB[1] + adjB[0], -adjB[0]}, dir
	case 3:
		return []int{adjB[1], -(adjB[0] + adjB[1])}, dir
	case 4:
		return []int{-adjB[0], -adjB[1]}, dir
	case 5:
		return []int{-(adjB[1] + adjB[0]), adjB[0]}, dir
	default:
		return []int{-adjB[1], adjB[0] + adjB[1]}, dir
	}
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
