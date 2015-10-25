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

func TestRingPolar(t *testing.T) {
	a := Coord{}
	r := rand.Intn(4) + 2
	ring := a.Ring(r)
	fmt.Println("RING POLAR TEST:")
	for _, test := range ring {
		fmt.Print(test.String() + "|" + test.Polar().String() + "|" + test.Polar().Coord().String() + "||")
	}
	fmt.Println("")
}

func TestRing(t *testing.T) {
	var fails int
	for i := 0; i < 100; i++ {
		a, b := GenTest()
		d := a.StepsTo(b)
		path := a.Ring(d)
		for i, step := range path {
			if a.StepsTo(step) != d {
				fmt.Print("ping1|", a, step, "|", a.StepsTo(step), d, "|")
				fails++
				panic("FAIL")
			}
			if i != 0 {
				if path[i-1].StepsTo(step) != 1 {
					fmt.Print("ping2")
					fails++
				}
			}
			if i != len(path)-1 {
				if path[i+1].StepsTo(step) != 1 {
					fmt.Print("ping3")
					fails++
				}
			}
			if i != len(path)-1 && i != 0 {
				if path[i-1] == path[i+1] {
					fmt.Print("ping4")
					fails++
				}
			}
		}
	}
	fmt.Println("Ring TEST fails:", fails)
}

func XTestPath(t *testing.T) {
	for i := 0; i < 10; i++ {
		a, b := GenTest()
		fmt.Println("------------------------------------------")
		fmt.Println("Testing:", a, b)
		d := a.StepsTo(b)
		fmt.Println("StepsTo:", d)
		p := a.PathTo(b)

		fmt.Println("D:", d, "PathD:", len(p), "Path:", p)
		if p[len(p)-1] != b || len(p) != d+1 {
			fmt.Println("============= FAILED =============")
		}
	}
}

func TestLotsPath(t *testing.T) {
	var fails int
	for i := 0; i < 1000; i++ {
		a, b := GenTest()
		d := a.StepsTo(b)
		p := a.PathTo(b)

		if p[len(p)-1] != b || len(p) != d+1 {
			fails++
		}
	}
	fmt.Println("BigTEST:", fails, "fails")
}

func xTestPixels(t *testing.T) {
	v := MakeViewport(20.0)
	v.CenterPix = Pixel{200, 200}
	for i := 0; i < 10; i++ {
		x := rand.Float64() * 200
		y := rand.Float64() * 200
		p := Pixel{x, y}
		hex := v.HexContaining(p)
		pt := v.CenterOf(hex)
		fmt.Println("Pixel", x, y, "is in hex", hex, "with center", pt[0], pt[1])
	}
}

func TestPixel(t *testing.T) {
	fmt.Println("-------------------------")
	v := MakeViewport(20.0)
	v.CenterPix = Pixel{200, 200}
	p := Pixel{215, 390}
	hex := v.HexContaining(p)
	pt := v.CenterOf(hex)
	fmt.Println("Pixel", p, "is in hex", hex, "with center", pt[0], pt[1])
	fmt.Println("-------------------------")
	p = Pixel{215, 410}
	hex = v.HexContaining(p)
	pt = v.CenterOf(hex)
	fmt.Println("Pixel", p, "is in hex", hex, "with center", pt[0], pt[1])
}

func GenTest() (a, b Coord) {
	return Coord{rand.Intn(100), rand.Intn(100)}, Coord{rand.Intn(100), rand.Intn(100)}
}
