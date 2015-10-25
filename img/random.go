package picture

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func bell5() int {
	switch rand.Intn(10) {
	case 0:
		return 1
	case 1, 2:
		return 2
	case 3, 4, 5, 6:
		return 3
	case 7, 8:
		return 4
	default:
		return 5
	}
}

func bell5Low() int {
	switch rand.Intn(6) {
	case 0:
		return 1
	case 1, 2:
		return 2
	default:
		return 3
	}
}

func bell5High() int {
	switch rand.Intn(6) {
	case 0:
		return 5
	case 1, 2:
		return 4
	default:
		return 3
	}
}

func wobble(n int) int {
	if coin() {
		n++
	}
	if coin() {
		n--
	}
	if n < 1 {
		return 1
	} else if n > 5 {
		return 5
	} else {
		return n
	}
}

func coin() bool {
	return rand.Intn(2) == 1
}

func pick(n int) int {
	return rand.Intn(n) + 1
}

func fpick(n int) float64 {
	return float64(rand.Intn(n) + 1)
}

func randWord(list []string) string {
	if list == nil {
		return ""
	}
	l := len(list)
	switch l {
	case 0:
		return ""
	case 1:
		return list[0]
	default:
		return list[rand.Intn(l)]
	}
}
