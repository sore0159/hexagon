package hexagon

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestFirst(t *testing.T) {
	fmt.Println("TESTING")
}

func TestPath(t *testing.T) {
	for i := 0; i < 10; i++ {
		a, b := GenTest()
		fmt.Println("------------------------------------------")
		fmt.Println("Testing:", a, b)
		d := a.StepsTo(b)
		p := a.PathTo(b)

		fmt.Println("D:", d, "PathD:", len(p), "Path:", p)
		if p[len(p)-1] != b || len(p) != d+1 {
			fmt.Println("============= FAILED =============")
		}
	}
}

func GenTest() (a, b Coord) {
	return Coord{rand.Intn(100), rand.Intn(100)}, Coord{rand.Intn(100), rand.Intn(100)}
}
