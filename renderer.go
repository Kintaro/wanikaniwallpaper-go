package main

import (
	"os"
	"flag"
	"bufio"
	"io/ioutil"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"code.google.com/p/freetype-go/freetype"
	"code.google.com/p/freetype-go/freetype/truetype"
)

var (
	dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	fontfile = flag.String("fontfile", "font/ipag.ttf", "filename of the ttf font")
	size     = flag.Float64("size", 12, "font size in points")
	width 	 = flag.Int("width", 1920, "width of wallpaper")
	height 	 = flag.Int("height", 1200, "height of wallpaper")
)

type Renderer struct {
	context *freetype.Context
	img draw.Image
	font *truetype.Font
	fontsize float64
}

func NewRenderer() *Renderer {
	r := Renderer {  }
	fontBytes, _ := ioutil.ReadFile(*fontfile)
	r.font, _ = freetype.ParseFont(fontBytes)
	r.img = image.NewRGBA(image.Rect(0, 0, *width, *height))
	r.context = freetype.NewContext()
	r.context.SetDPI(*dpi)
	r.context.SetFont(r.font)
	r.context.SetFontSize(*size)
	r.context.SetClip(r.img.Bounds())
	draw.Draw(r.img, r.img.Bounds(), &image.Uniform{ color.RGBA{ 0, 0, 0, 255 } }, image.ZP, draw.Src)
	r.context.SetDst(r.img)

	return &r
}

func (r *Renderer) SetFontSize(size int) {
	r.context.SetFontSize(float64(size))
	r.fontsize = float64(size)
}

func (r *Renderer) DrawKanji(kanji *Kanji, x int, y int) {
	pt := freetype.Pt(x, y - int(r.fontsize / 10.0))
	pt.Y += r.context.PointToFix32(r.fontsize)
	r.context.SetSrc(&image.Uniform{ kanji.Color() })

	for _, s := range kanji.character {
		r.context.DrawString(string(s), pt)
	}
	r.context.SetClip(r.img.Bounds())
}

func (r *Renderer) SaveImage(filename string) {
	file, _ := os.Create(filename)
	defer file.Close()

	b := bufio.NewWriter(file)
	png.Encode(b, r.img)
	b.Flush()
}
