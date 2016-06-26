// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"sync"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	outchan := make(chan color.Color, 8)
	var waitgroup sync.WaitGroup

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			waitgroup.Add(1)
			go mandelbrot(z, outchan, &waitgroup)
			// Image point (px, py) represents complex value z.

			img.Set(px, py, <- outchan)
		}
	}
	waitgroup.Wait()
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128, outchan chan color.Color, waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()

	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			outchan <- color.Gray{255 - contrast*n}
			return
		}
	}
	outchan <- color.Black
	return
}
