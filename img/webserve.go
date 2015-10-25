package picture

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"net/http"
)

func ServePic(w http.ResponseWriter) {
	dest := image.NewRGBA(image.Rect(0, 0, 150, 400))
	grey := color.RGBA{0x22, 0x22, 0x22, 255}
	draw.Draw(dest, dest.Bounds(), &image.Uniform{grey}, image.ZP, draw.Src)
	gc := draw2dimg.NewGraphicContext(dest)
	_ = gc
	err := png.Encode(w, dest)
	if err != nil {
		log.Println(err)
	}
}
