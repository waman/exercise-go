package main

import (
	"image"
	"image/png"
	"os"
	"image/color"
	"math/cmplx"
)

const(
	xmin, xmax    = -0.5, 0
	ymin, ymax    = -1, -0.5
	width, height = 512, 512
)

func main(){
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrotC128(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrotC128(z complex128) color.Color {
	const iterations = 200
	const contrast   = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{ Y:255 - contrast*n }
		}
	}
	return color.Black
}
