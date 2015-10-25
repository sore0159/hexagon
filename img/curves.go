package picture

import (
	//	"fmt"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dkit"
	"image/color"
)

func Line(gc draw2d.GraphicContext, p1, p2 p) {
	gc.MoveTo(p1[0], p1[1])
	gc.LineTo(p2[0], p2[1])
	gc.Stroke()
}
func Blip(gc draw2d.GraphicContext, point p) {
	Circle(gc, [2]float64{point[0], point[1]}, 5)
	gc.FillStroke()
}
func Blotch(gc draw2d.GraphicContext, point p, n int) {
	gc.SetLineWidth(1)
	gc.SetStrokeColor(Lighter(Pallate("white"), .50))
	gc.SetFillColor(color.RGBA{0xCC, 0xCC, 0xCC, 0xFF})
	p1 := point.add(p{fpick(5), 0, 10})
	p2 := point.add(p{fpick(5), fpick(5), 10})
	p3 := point.add(p{0, fpick(5), 10})
	p4 := point.add(p{-fpick(5), 0, 10})
	p5 := point.add(p{-fpick(5), -fpick(5), 10})
	p6 := point.add(p{0, -fpick(5), 10})

	path := []p{p1, p2, p3, p4, p5, p6}
	RoundPoly(gc, path...)
}

func RoundPoly(path draw2d.GraphicContext, points ...p) {
	if len(points) < 3 {
		return
	}
	for i, point := range points {
		var nextPoint, prevPoint p
		if i == len(points)-1 {
			nextPoint = points[0]
		} else {
			nextPoint = points[i+1]
		}
		if i == 0 {
			prevPoint = points[len(points)-1]
		} else {
			prevPoint = points[i-1]
		}
		nextMid := midPoint(point, nextPoint)
		prevMid := midPoint(point, prevPoint)
		if i == 0 {
			path.MoveTo(prevMid[0], prevMid[1])
		}
		curve := point[2]
		here := [2]float64{point[0], point[1]}
		BendPath(path, prevMid, nextMid, here, curve)
	}
	path.FillStroke()
}

func RoundPath(path draw2d.GraphicContext, points ...p) {
	if len(points) < 3 {
		return
	}
	for i, point := range points {
		var nextMid, prevMid [2]float64
		curve := point[2]
		here := [2]float64{point[0], point[1]}
		if i == len(points)-1 {
		} else {
			nextMid = midPoint(point, points[i+1])
		}
		if i == 0 {
			path.LineTo(nextMid[0], nextMid[1])
		} else {
			prevMid = midPoint(point, points[i-1])
			if i == len(points)-1 {
				path.LineTo(here[0], here[1])
			} else {
				BendPath(path, prevMid, nextMid, here, curve)
			}
		}
	}
}

func BendPath(path draw2d.PathBuilder, a, b, bend [2]float64, curve float64) {
	if a == b {
		//fmt.Println("ping1")
		return
	}
	path.MoveTo(a[0], a[1])
	if curve == 0 {
		path.LineTo(bend[0], bend[1])
		path.LineTo(b[0], b[1])
		//fmt.Println("ping2")
		return
	}
	dA := dist(a, bend)
	dB := dist(b, bend)
	if dA < curve {
		curve = dA
	}
	if dB < curve {
		curve = dB
	}
	if bend[0]-a[0] == 0 || b[0]-bend[0] == 0 {
		if bend[0]-a[0] == 0 && b[0]-bend[0] == 0 {
			path.LineTo(b[0], b[1])
			//fmt.Println("ping3")
			return
		}
	} else {
		slopeA := (bend[1] - a[1]) / (bend[0] - a[0])
		slopeB := (b[1] - bend[1]) / (b[0] - bend[0])
		if slopeA == slopeB {
			path.LineTo(b[0], b[1])
			//fmt.Println("ping4")
			return
		}
	}
	//fmt.Println("DA, DB:", dA, dB)
	stop1 := gotwoard(a, bend, dA-curve)
	stop2 := gotwoard(b, bend, dB-curve)
	/*fmt.Println("")
	fmt.Println("CURVE:", curve)
	fmt.Println(a, "TO", stop1)
	fmt.Println(stop1, "CURVE TO", stop2, "AROUND", bend)
	fmt.Println(stop2, "TO", b)
	*/
	path.LineTo(stop1[0], stop1[1])
	path.QuadCurveTo(bend[0], bend[1], stop2[0], stop2[1])
	path.LineTo(b[0], b[1])
}

func RoundedCurve(path draw2d.PathBuilder, x1, y1, x2, y2, topArc, botArc float64) {
	var arc1, arc2 float64
	if y2 > y1 {
		arc1, arc2 = topArc, botArc
	} else {
		arc1, arc2 = botArc, topArc
	}
	if arc1 == 0 {
		arc1 = 1
	}
	if arc2 == 0 {
		arc2 = 1
	}
	botX1 := x1 + 10
	botX2 := x2 - 10

	path.MoveTo(x1, y1+arc1)
	path.QuadCurveTo(x1, y1, x1+arc1, y1)
	path.LineTo(x2-arc1, y1)
	path.QuadCurveTo(x2, y1, x2, y1+arc1)
	path.LineTo(botX2, y2-arc2)
	path.QuadCurveTo(botX2, y2, botX2-arc2, y2)
	path.LineTo(botX1+arc2, y2)
	path.QuadCurveTo(botX1, y2, botX1, y2-arc2)
	path.Close()
}

func Circle(path draw2d.PathBuilder, center [2]float64, rad float64) {
	draw2dkit.Circle(path, center[0], center[1], rad)
}

func Oval(path draw2d.PathBuilder, center [2]float64, rad float64) {
}
