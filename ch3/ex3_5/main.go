// 【練習問題 3.5】
// image.NewRGBA と color.RGBA 型あるいは color.YCbCr 型を使って
// フルカラーのマンデルブロ集合を実装しなさい。
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"os"
)

func main() {
	var w io.Writer

	if len(os.Args) <= 1 {
		w = os.Stdout
	} else {
		// 引数があればファイルに出力（os パッケージのドキュメント参照）
		file, err := os.OpenFile(os.Args[1], os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0755)
		if err != nil {
			log.Fatal(err)
		}

		defer func() {
			if cErr := file.Close(); err == nil && cErr != nil {
				log.Fatal(err)
			}
		}()

		w = file
	}

	outputImage(w)
}

func outputImage(w io.Writer) {
	const (
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
	png.Encode(w, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 255
	const contrast = 30

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			y := 255 - contrast*n
			return color.RGBA{y, 0xff, y, 0xff}
		}
	}
	return color.RGBA{0x00, 0x00, 0xff, 0xff}
}

// palette.Plan9 を使ったバージョン
//func mandelbrot(z complex128) color.Color {
//	const iterations = 255
//	const contrast   = 31
//
//	var v complex128
//	for n := 0; n < iterations; n++ {
//		v = v*v + z
//		if cmplx.Abs(v) > 2 {
//			// 配色を考えるのが面倒なので、標準パッケージに定義済みの
//			// palette.Plan9 を使用（問題文無視）。
//			return palette.Plan9[n*contrast % 256]
//		}
//	}
//	return color.White
//}
