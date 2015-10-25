package picture

import (
	//	"fmt"
	"math"
)

type p [3]float64

func RefX(path []p, x float64) []p {
	r := make([]p, len(path))
	for i, pt := range path {
		r[len(r)-i-1] = pt.refX(x)
	}
	return r
}

func Points(path []p) []p {
	points := make([]p, len(path))
	for i, pt := range path {
		if i == 0 {
			points[0] = pt
		} else {
			points[i] = points[i-1].add(pt)
		}
	}
	return points
}

func (a p) refX(x float64) p {
	b := a
	b[0] = 2*x - a[0]
	return b
}
func (a p) add(b p) p {
	return p{a[0] + b[0], a[1] + b[1], b[2]}
}

func (a p) slope(b p) float64 {
	return (b[1] - a[1]) / (b[0] - a[0])
}

func midPoint(a, b p) [2]float64 {
	c := [2]float64{(b[0] + a[0]) / 2, (a[1] + b[1]) / 2}
	//fmt.Println("MID:", a, b, c)
	return c
}

func dist(a, b [2]float64) float64 {
	return math.Sqrt(math.Pow(a[0]-b[0], 2) + math.Pow(a[1]-b[1], 2))
}
func gotwoard(a, b [2]float64, dist float64) [2]float64 {
	if dist == 0 {
		return a
	}
	if b[0] == a[0] {
		c := a
		if a[1] > b[1] {
			c[1] = c[1] - dist
		} else {
			c[1] = c[1] + dist
		}
		return c
	}
	slope := (b[1] - a[1]) / (b[0] - a[0])
	// d^2 = ax-cx)^2 + (ay-cy)^2
	// d^2 = ax-cx)^2 + (slope*(ax-cx))^2
	c := [2]float64{}
	if a[0] < b[0] {
		c[0] = a[0] + math.Sqrt((dist*dist)/(1+slope*slope))
	} else {
		c[0] = a[0] - math.Sqrt((dist*dist)/(1+slope*slope))
	}
	// slope = (c[1] - a[1])/(c[0] - a[0])
	c[1] = a[1] + slope*(c[0]-a[0])
	return c
}
