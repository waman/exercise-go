// 【練習問題 3.5】
// image.NewRGBA と color.RGBA 型あるいは color.YCbCr 型を使って
// フルカラーのマンデルブロ集合を実装しなさい。
package main

import (
	"image"
	"image/png"
	"os"
	"image/color"
	"math/cmplx"
)

func main(){
	const(
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 255
	const contrast   = 30

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			y := 255 - contrast*n
			return color.RGBA{ y, 0xff, y, 0xff }
		}
	}
	return color.RGBA{ 0x00, 0x00, 0xff, 0xff }
}
