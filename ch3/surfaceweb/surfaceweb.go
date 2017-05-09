package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
)

var xyrange = 30.0                // axis ranges (-xyrange..+xyrange)
var xyscale = width / 2 / xyrange // pixels per x or y unit

const (
	width, height = 600, 320     // canvas size in pixels
	cells         = 100          // number of grid cells
	zscale        = height * 0.4 // pixels per z unit
	angle         = math.Pi / 6  // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/surface", handlerSurface)
	err := http.ListenAndServe("localhost:8200", nil)
	log.Fatal(err)
}

func generateSurface() string {
	xyrange = rand.Float64() * 100.
	xyscale = width / 2 / xyrange
	surface := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			surface += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	surface += "</svg>"
	return surface
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Please go to http://localhost:8200/surface to see the svg")
}

func handlerSurface(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintf(w, generateSurface())
}

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

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
