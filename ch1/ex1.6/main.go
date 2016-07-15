/*
Exercise 1.6:
Modify the Lissajous program to produce images in multiple colors by adding more
values to palette and then displaying them by changing the third argument of
SetColorIndex in some interesting way.

Solution:
Change the front ground color every frame by changing palette index.
*/
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var red = color.RGBA{0xff, 0, 0, 0xff}
var green = color.RGBA{0, 0xff, 0, 0xff}
var blue = color.RGBA{0, 0, 0xff, 0xff}
var palette = []color.Color{color.Black, red, green, blue}

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas conver [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5),
				size+int(y*size+0.5),
				uint8(i%3+1))

		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
