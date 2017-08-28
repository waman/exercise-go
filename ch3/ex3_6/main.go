// 【練習問題 3.6】
// スーパーサンプリング (supersampling) は、個々の画素内の複数の点のカラー値を
// 計算して平均を求めることでピクセル化の影響を薄める技法です。最も単純な方法は、
// 個々の画素を四つの「サブピクセル」へ分割することです。その方法を実装しなさい。
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
		dx = float64(xmax-xmin)/width  // 画素の x 方向の幅（もしくは隣り合う画素の左上端どうしの距離）
		dy = float64(ymax-ymin)/height // 画素の y 方向の幅
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)*dy + ymin
		for px := 0; px < width; px++ {
			x := float64(px)*dx + xmin

			// 1つの画素を4つのサブピクセルに分割したときの、
			// サブピクセルの各左上端について色を計算。
			c1 := mandelbrot(complex(x, y))
			c2 := mandelbrot(complex(x+dx/2, y))
			c3 := mandelbrot(complex(x, y+dy/2))
			c4 := mandelbrot(complex(x+dx/2, y+dy/2))

			// 色の平均。桁あふれを起こすので大きい型に変換してから計算している。
			ave := uint8((uint16(c1.Y)+uint16(c2.Y)+uint16(c3.Y)+uint16(c4.Y))/4)
			c := color.Gray{ ave }

			img.Set(px, py, c)
		}
	}
	png.Encode(os.Stdout, img)
}

// 色の平均を計算しやすくするため、返り値を Gray にしています。
// 繰り返しを抜けた場合の黒色も Gray{0} に変更しています。
func mandelbrot(z complex128) color.Gray {
	const iterations = 200
	const contrast   = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Gray{0}  // 黒色
}