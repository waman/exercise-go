package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"os"
)

func main() {
	var w io.Writer

	if len(os.Args) <= 1 {
		w = os.Stdout
	} else {
		// 引数があればファイルに出力（os パッケージのドキュメント参照）
		file, err := os.Create(os.Args[1])
		if err != nil { log.Fatal(err) }

		defer func() {
			if err := file.Close(); err != nil { log.Fatal(err)	}
		}()

		w = file
	}

	outputImageC64(w)
}

func outputImageC64(w io.Writer) {
	const (
		xmin, xmax    = -0.5, 0
		ymin, ymax    = -1, -0.5
		width, height = 512, 512
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float32(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float32(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrotC64(z))
		}
	}
	png.Encode(w, img)
}

func mandelbrotC64(z complex64) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		re, im := real(v), imag(v)
		if re*re*im*im > 4 {
			return color.Gray{Y: 255 - contrast*n}
		}
	}
	return color.Black
}
