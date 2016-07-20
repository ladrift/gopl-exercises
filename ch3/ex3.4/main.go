/*
Exercise 3.4:
Following the approac of the Lissajous example in Section 1.7, constructa web server
that computes surfaces and writes SVG data to the client. The server must set the
Content-Type header like this:
	w.Header().Set("Content-Type", "image/svg+xml")
Allow the client to specify values like height, width, and color as HTTP request
parameters.
Solution:
Currently, only support color query parameter.
*/
package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "image/svg+xml")
			if err := r.ParseForm(); err != nil {
				fmt.Fprintln(os.Stderr, "ParseForm failed:", err)
			}
			color := r.Form.Get("color")
			if color == "" {
				color = "white"
			}
			writeSVG(w, color)
		}
		http.HandleFunc("/", handler)

		log.Fatal(http.ListenAndServe("localhost:8000", nil))
	}
	writeSVG(os.Stdout, "white")
}

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func writeSVG(out io.Writer, color string) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: %s; stroke-width: 0.7' "+
		"width='%d' height='%d'>", color, width, height)
	var err error
	cornerErr := func(i int, j int) (float64, float64) {
		x, y := corner(i, j)
		if math.IsNaN(x) || math.IsNaN(y) {
			err = errors.New("corner return NaN")
		}
		return x, y
	}
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			err = nil
			ax, ay := cornerErr(i+1, j)
			bx, by := cornerErr(i, j)
			cx, cy := cornerErr(i, j+1)
			dx, dy := cornerErr(i+1, j+1)
			if err != nil {
				// Skip this polygon
				continue
			}
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

var f = bessel

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func bessel(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.J0(r)
}
