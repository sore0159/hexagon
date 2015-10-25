package picture

import (
	"fmt"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	//	"image/color"
	"image/draw"
	"log"
	"mule/hexagon"
	"testing"
)

func TestFirst(t *testing.T) {
	log.Println("TESTING")
}

func TestSecond(t *testing.T) {
	draw2d.SetFontFolder("")
	dest := image.NewRGBA(image.Rect(0, 0, 400, 400))
	white := Pallate("white")
	draw.Draw(dest, dest.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)
	gc := draw2dimg.NewGraphicContext(dest)
	gc.SetFontData(draw2d.FontData{Name: "DroidSansMono", Family: draw2d.FontFamilyMono})
	gc.SetFillColor(image.Black)
	gc.SetFontSize(8)
	v := GetVP()
	for x := 0; x < 400; x++ {
		for y := 0; y < 400; y++ {
			p := hexagon.Pixel{float64(x), float64(y)}
			hex := v.HexContaining(p)
			//hex := v.HexContaining(p.Flip())
			clr := Pallate("blue")
			clN := (hex[0] - hex[1]) % 4
			for clN < 0 {
				clN += 4
			}
			switch clN {
			case 0:
				clr = Pallate("blue")
			case 1:
				clr = Pallate("green")
			case 2:
				clr = Pallate("red")
			default:
				clr = Pallate("white")
			}
			_ = clr
			dest.Set(x, y, clr)
			/*if (x+60)%60 == 0 && y%15 == 0 {
				//pt := v.CenterOf(hex)
				//fmt.Println("Pixel", p, "is in hex", hex, "with center", pt[0], pt[1], "and color", (hex[0]+hex[1])%4)
				gc.FillStringAt(fmt.Sprintf("(%d,%d)", hex[0], hex[1]), p[0], p[1])
			}*/
		}
	}
	for x := -5; x < 5; x++ {
		for y := -5; y < 5; y++ {
			//for x := 0; x < 2; x++ {
			//for y := 0; y < 2; y++ {
			h := v.CenterHex.Add(hexagon.Coord{x, y})
			corners := v.CornersOf(h)
			/*for i, cn := range corners {
				corners[i] = cn.Flip()
			}*/
			gc.MoveTo(corners[0][0], corners[0][1])
			for i, _ := range corners {
				var px, py float64
				if i == 5 {
					px, py = corners[0][0], corners[0][1]
				} else {
					px, py = corners[i+1][0], corners[i+1][1]
				}
				gc.LineTo(px, py)
			}
			gc.Close()
			gc.Stroke()
			c := v.CenterOf(h)
			//c := v.CenterOf(h).Flip()
			gc.FillStringAt(fmt.Sprintf("(%d,%d)", h[0], h[1]), c[0], c[1])
		}
	}

	draw2dimg.SaveToPngFile("picture.png", dest)
}

func GetVP() *hexagon.Viewport {
	v := hexagon.MakeViewport(40.0)
	v.CenterPix = hexagon.Pixel{100, 200}
	v.CenterHex = hexagon.Coord{10, 10}
	//v.HexH *= .75 // For isometric slant; it's that easy!
	//v.HexW *= .75 // For flipped isometric slant; it's that easy!
	return v
}
