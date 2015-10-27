package hexagon

import (
	"fmt"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
	"image/draw"
	"log"
	"testing"
)

func TestFirstPIC(t *testing.T) {
	log.Println("TESTING")
}

func TestSecondPIC(t *testing.T) {
	v := GetVP()
	v.SetFlattop(true)
	v.SetIso(true)
	v.SetIso(false)
	v.SetFlattop(false)
	v.SetIso(true)
	DrawGrid(v, "fliptop.png")
}

func TestThirdPIC(t *testing.T) {
	v := GetVP()
	v.SetFlattop(true)
	v.SetFlattop(false)
	v.SetFlattop(true)
	v.SetIso(true)
	v.SetIso(false)
	v.SetIso(true)
	DrawGrid(v, "flattop.png")
}

func DrawGrid(v *Viewport, fileName string) {
	draw2d.SetFontFolder("")
	dest := image.NewRGBA(image.Rect(0, 0, 400, 400))
	white := color.RGBA{0xff, 0xff, 0xff, 255}
	draw.Draw(dest, dest.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)
	gc := draw2dimg.NewGraphicContext(dest)
	gc.SetFontData(draw2d.FontData{Name: "DroidSansMono", Family: draw2d.FontFamilyMono})
	gc.SetFillColor(image.Black)
	gc.SetFontSize(8)
	for x := 0; x < 400; x++ {
		for y := 0; y < 400; y++ {
			p := Pixel{float64(x), float64(y)}
			hex := v.HexContaining(p)
			var clr color.RGBA
			clN := (hex[0] - hex[1]) % 4
			for clN < 0 {
				clN += 4
			}
			switch clN {
			case 0: // blue
				clr = color.RGBA{0x33, 0x66, 0xcc, 255}
			case 1: // green
				clr = color.RGBA{0x00, 0xAA, 0x33, 255}
			case 2: // red
				clr = color.RGBA{0xAA, 0x33, 0x33, 255}
			default: // white
				clr = white
			}
			dest.Set(x, y, clr)
			/*if (x+60)%60 == 0 && y%15 == 0 {
				gc.FillStringAt(fmt.Sprintf("(%d,%d)", hex[0], hex[1]), p[0], p[1])
			}*/
		}
	}
	for _, h := range v.VisList() {
		corners := v.CornersOf(h)
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
		gc.FillStringAt(fmt.Sprintf("(%d,%d)", h[0], h[1]), c[0], c[1])
	}
	gc.SetLineWidth(2)
	gc.MoveTo(v.ULCorner[0], v.ULCorner[1])
	gc.LineTo(v.LRCorner[0], v.ULCorner[1])
	gc.LineTo(v.LRCorner[0], v.LRCorner[1])
	gc.LineTo(v.ULCorner[0], v.LRCorner[1])
	gc.Close()
	gc.Stroke()

	//}

	draw2dimg.SaveToPngFile(fileName, dest)
}

func GetVP() *Viewport {
	v := MakeViewport(30.0, false, false)
	v.CenterPix = Pixel{200, 200}
	v.CenterHex = Coord{1, 5}
	v.SetFrame(Pixel{115, 85}, Pixel{305, 315})
	return v
}
