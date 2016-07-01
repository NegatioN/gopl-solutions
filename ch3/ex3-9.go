package main

/*
Write a web server that renders fractals and writes the image data tot he client.
Allow the client to specify the x,y, and zoom values as parameters to the HTTP request.
 */

import (
	"log"
	"net/http"
	"sync"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"strconv"
)

var mu sync.Mutex

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}


// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	width, werr := strconv.ParseFloat(r.FormValue("width"), 64)
	height, herr := strconv.ParseFloat(r.FormValue("height"), 64)
	zoom, zerr := strconv.ParseFloat(r.FormValue("zoom"), 64)

	if err != nil || werr != nil || herr != nil || zerr != nil{
		makeUserPng(100, 100, 2, w)
	}

	makeUserPng(width,height,zoom, w)
}



//DEFAULT methods. Must implement zoomable mandlebrot yourself. Not in scope of 3.9
func makeUserPng(width float64, height float64, zoom float64, w http.ResponseWriter) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
	)
	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	for py := 0; py < int(height); py++ {
		y := float64(py)/ height*(ymax-ymin) + ymin
		for px := 0; px < int(width); px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z, zoom))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func mandelbrot(zCord complex128, zoom float64) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + zCord
		if cmplx.Abs(v) > zoom {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}