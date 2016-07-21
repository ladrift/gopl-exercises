// Exercise 3.6:
// Supersampling is a technique to reduce the effect of pixelation by computing
// the color value at several points within each pixel and taking the average.
// The simplest method is to divide each pixel into four "subpixel".
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		y1 := y + (1/4.0)/height*(ymax-ymin)
		y2 := y + (3/4.0)/height*(ymax-ymin)
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			x1 := x + (1/4.0)/height*(xmax-xmin)
			x2 := x + (3/4.0)/height*(xmax-xmin)
			z1 := complex(x1, y1)
			z2 := complex(x1, y2)
			z3 := complex(x2, y1)
			z4 := complex(x2, y2)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrotAve(z1, z2, z3, z4))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrotAve(zs ...complex128) color.Color {
	count := uint8(len(zs))
	var r, g, b, a uint8
	for _, z := range zs {
		c := mandelbrot(z)
		rt, gt, bt, at := c.RGBA()
		r += uint8(rt) / count
		g += uint8(gt) / count
		b += uint8(bt) / count
		a += uint8(at) / count
	}
	return color.RGBA{r, g, b, a}

}
func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
